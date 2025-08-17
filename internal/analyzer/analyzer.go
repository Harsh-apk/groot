package analyzer

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/harsh-apk/groot/internal/model"
	"github.com/harsh-apk/groot/internal/parser"
	"github.com/harsh-apk/groot/internal/walker"
)

// GetLanguageByFileExtension finds the appropriate language configuration from the compiled list.
func GetLanguageByFileExtension(filePath string) (model.Language, bool) {
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext == "" {
		return model.Language{}, false
	}
	for _, lang := range CompiledLanguageConfig.Languages {
		for _, langExt := range lang.FileExtensions {
			if ext == langExt {
				return lang, true
			}
		}
	}
	return model.Language{}, false
}

// Analyze performs the core analysis and returns the raw data structures.
func Analyze(rootPath string, skipDirs []string, includeExts []string) (*model.Node, model.Analytics, error) {
	startTime := time.Now()

	rootNode, err := walker.BuildFileTree(rootPath, skipDirs)
	if err != nil {
		return nil, model.Analytics{}, fmt.Errorf("failed to walk file tree: %w", err)
	}

	allFileNodes := collectFileNodes(rootNode)
	var filteredFileNodes []*model.Node

	if len(includeExts) > 0 {
		includeSet := make(map[string]struct{})
		for _, ext := range includeExts {
			includeSet[strings.TrimSpace(ext)] = struct{}{}
		}
		for _, node := range allFileNodes {
			ext := filepath.Ext(node.Path)
			if _, ok := includeSet[ext]; ok {
				filteredFileNodes = append(filteredFileNodes, node)
			}
		}
	} else {
		filteredFileNodes = allFileNodes
	}

	var wg sync.WaitGroup
	jobs := make(chan *model.Node, len(filteredFileNodes))
	for w := 0; w < runtime.NumCPU(); w++ {
		wg.Add(1)
		go worker(&wg, jobs)
	}
	for _, node := range filteredFileNodes {
		jobs <- node
	}
	close(jobs)
	wg.Wait()

	stats := aggregateAnalytics(allFileNodes, filteredFileNodes)
	stats.Duration = time.Since(startTime)
	stats.DurationReadable = stats.Duration.Round(time.Millisecond).String()
	stats.FilesScanned = len(allFileNodes)

	return rootNode, stats, nil
}

// FormatText takes the raw analysis data and generates the human-readable string outputs.
// Note: The calling function in cmd/analyze.go should be updated to pass 'includeExts'.
func FormatText(rootNode *model.Node, stats model.Analytics, includeExts []string) (string, string) {
	var treeBuilder strings.Builder
	absPath, _ := filepath.Abs(rootNode.Path)
	treeBuilder.WriteString(fmt.Sprintf("Codebase overview for: %s\n\n", absPath))
	formatTree(&treeBuilder, rootNode, "", true, includeExts)

	var analyticsBuilder strings.Builder
	appendAnalytics(&analyticsBuilder, stats)

	return treeBuilder.String(), analyticsBuilder.String()
}

// worker is a concurrent worker that parses file nodes.
func worker(wg *sync.WaitGroup, jobs <-chan *model.Node) {
	defer wg.Done()
	for node := range jobs {
		lang, supported := GetLanguageByFileExtension(node.Path)
		if !supported {
			continue
		}
		content, err := os.ReadFile(node.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not read file %s: %v\n", node.Path, err)
			continue
		}
		node.LOC = bytes.Count(content, []byte("\n")) + 1
		elements, err := parser.Parse(content, lang)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not parse file %s: %v\n", node.Path, err)
			continue
		}
		node.CodeElements = elements
	}
}

// aggregateAnalytics processes file nodes to build the analytics summary.
func aggregateAnalytics(allNodes, parsedNodes []*model.Node) model.Analytics {
	stats := model.Analytics{
		PerLanguageStats: make(map[string]model.LanguageStats),
	}
	for _, node := range parsedNodes {
		lang, supported := GetLanguageByFileExtension(node.Path)
		if !supported {
			continue
		}
		if node.LOC > 0 {
			stats.FilesParsed++
			stats.TotalLOC += node.LOC
		}
		stats.TotalElements += len(node.CodeElements)
		langStats, ok := stats.PerLanguageStats[lang.Name]
		if !ok {
			langStats = model.LanguageStats{ElementCounts: make(map[string]int)}
		}
		langStats.FileCount++
		langStats.LOC += node.LOC
		for _, el := range node.CodeElements {
			langStats.ElementCounts[el.Type]++
		}
		stats.PerLanguageStats[lang.Name] = langStats
	}
	return stats
}

// appendAnalytics formats and writes the analytics summary.
func appendAnalytics(builder *strings.Builder, stats model.Analytics) {
	builder.WriteString("\n\n---\n\n")
	builder.WriteString("📊 Analysis Report\n\n")
	builder.WriteString("Overall Summary\n")
	builder.WriteString("────────────────────────────────────────\n")
	builder.WriteString(fmt.Sprintf("%-20s %s\n", "Analysis Duration:", stats.DurationReadable))
	builder.WriteString(fmt.Sprintf("%-20s %d\n", "Files Scanned:", stats.FilesScanned))
	builder.WriteString(fmt.Sprintf("%-20s %d\n", "Files Parsed:", stats.FilesParsed))
	builder.WriteString(fmt.Sprintf("%-20s %d\n", "Total Lines of Code:", stats.TotalLOC))
	builder.WriteString(fmt.Sprintf("%-20s %d\n", "Total Elements Found:", stats.TotalElements))
	builder.WriteString("────────────────────────────────────────\n\n")
	builder.WriteString("Language Breakdown\n")
	builder.WriteString("────────────────────────────────────────\n")

	sortedLangs := make([]string, 0, len(stats.PerLanguageStats))
	for langName := range stats.PerLanguageStats {
		sortedLangs = append(sortedLangs, langName)
	}
	sort.Strings(sortedLangs)

	for _, langName := range sortedLangs {
		langStats := stats.PerLanguageStats[langName]
		builder.WriteString(fmt.Sprintf("▶ %s (%d files, %d LOC)\n", langName, langStats.FileCount, langStats.LOC))
		if len(langStats.ElementCounts) > 0 {
			sortedElements := make([]string, 0, len(langStats.ElementCounts))
			for elType := range langStats.ElementCounts {
				sortedElements = append(sortedElements, elType)
			}
			sort.Strings(sortedElements)
			for _, elType := range sortedElements {
				builder.WriteString(fmt.Sprintf("  - %-18s %d\n", elType+":", langStats.ElementCounts[elType]))
			}
		}
		builder.WriteString("\n")
	}
}

// collectFileNodes recursively traverses the tree and returns a flat slice of all file nodes.
func collectFileNodes(node *model.Node) []*model.Node {
	var files []*model.Node
	if !node.IsDir {
		files = append(files, node)
		return files
	}
	for _, child := range node.Children {
		files = append(files, collectFileNodes(child)...)
	}
	return files
}

// formatTree recursively builds the string representation of the file tree.
// It now filters the display based on the includeExts list.
func formatTree(builder *strings.Builder, node *model.Node, prefix string, isRoot bool, includeExts []string) {
	// If a filter is active, only show directories that contain included files.
	if node.IsDir && len(includeExts) > 0 && !isRoot {
		if !directoryContainsIncludedFiles(node, includeExts) {
			return
		}
	}

	name := filepath.Base(node.Path)
	if isRoot {
		name = node.Path
	}
	builder.WriteString(name + "\n")

	if !node.IsDir && len(node.CodeElements) > 0 {
		sort.Slice(node.CodeElements, func(i, j int) bool {
			return node.CodeElements[i].Line < node.CodeElements[j].Line
		})
		for _, el := range node.CodeElements {
			builder.WriteString(fmt.Sprintf("%s  - %s: %s (L%d)\n", prefix, el.Type, el.Name, el.Line))
		}
	}

	for i, child := range node.Children {
		isLast := i == len(node.Children)-1
		connector := "├── "
		newPrefix := prefix + "│   "
		if isLast {
			connector = "└── "
			newPrefix = prefix + "    "
		}

		// If the child is a directory, recurse into it.
		if child.IsDir {
			builder.WriteString(prefix + connector)
			formatTree(builder, child, newPrefix, false, includeExts)
		} else if len(includeExts) == 0 {
			// If no filter, print all files.
			builder.WriteString(prefix + connector)
			formatTree(builder, child, newPrefix, false, includeExts)
		} else {
			// If filtering, only print files that are in the include list.
			isIncluded := false
			ext := filepath.Ext(child.Path)
			for _, includedExt := range includeExts {
				if ext == includedExt {
					isIncluded = true
					break
				}
			}
			if isIncluded {
				builder.WriteString(prefix + connector)
				formatTree(builder, child, newPrefix, false, includeExts)
			}
		}
	}
}

// directoryContainsIncludedFiles is a helper to check if a directory or any of its
// subdirectories contain a file that matches the include list.
func directoryContainsIncludedFiles(node *model.Node, includeExts []string) bool {
	if !node.IsDir {
		return false
	}

	for _, child := range node.Children {
		if !child.IsDir {
			ext := filepath.Ext(child.Path)
			for _, includedExt := range includeExts {
				if ext == includedExt {
					return true // Found a matching file
				}
			}
		} else {
			// Recurse into subdirectory
			if directoryContainsIncludedFiles(child, includeExts) {
				return true
			}
		}
	}
	return false
}
