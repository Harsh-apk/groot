package walker

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/denormal/go-gitignore"
	"github.com/harsh-apk/groot/internal/model"
)

// defaultIgnorePatterns specifies a comprehensive list of directories and file patterns
// that are almost always excluded from a codebase analysis.
var defaultIgnorePatterns = []string{
	// --- General & OS-specific ---
	".DS_Store",
	"*.log",
	"*.lock",

	// --- Version Control ---
	".git",

	// --- Editor & IDE Config ---
	".vscode",
	".idea",

	// --- JavaScript / Node.js ---
	"node_modules",
	".next",
	"dist",
	"build",
	"coverage",

	// --- Go ---
	"vendor",

	// --- Python ---
	"__pycache__",
	".venv",
	"venv",
	"*.pyc",
	".env",

	// --- Java / Maven / Gradle ---
	"target",
	".gradle",

	// --- Ruby ---
	"tmp",
	"log",

	// --- Elixir ---
	"_build",
	"deps",
}

// --- UPDATED FUNCTION SIGNATURE ---
// BuildFileTree now accepts custom ignore patterns from the user.
func BuildFileTree(rootPath string, customIgnorePatterns []string) (*model.Node, error) {
	absRoot, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, fmt.Errorf("could not get absolute path for '%s': %w", rootPath, err)
	}

	// --- COMBINE IGNORE PATTERNS ---
	// Start with the comprehensive default ignore patterns.
	patterns := defaultIgnorePatterns
	// Append any custom patterns provided by the user via the --skip flag.
	patterns = append(patterns, customIgnorePatterns...)

	// Read and append patterns from the project's .gitignore file if it exists.
	ignoreFilePath := filepath.Join(absRoot, ".gitignore")
	if _, err := os.Stat(ignoreFilePath); err == nil {
		content, err := os.ReadFile(ignoreFilePath)
		if err == nil {
			patterns = append(patterns, strings.Split(string(content), "\n")...)
		}
	}

	// Create a gitignore matcher with the combined patterns.
	ignore := gitignore.New(
		strings.NewReader(strings.Join(patterns, "\n")),
		absRoot,
		nil,
	)

	rootNode := &model.Node{Name: filepath.Base(absRoot), Path: absRoot, IsDir: true}
	nodesByPath := map[string]*model.Node{absRoot: rootNode}

	walkErr := filepath.WalkDir(absRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == absRoot {
			return nil
		}

		if match := ignore.Match(path); match != nil && match.Ignore() {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		node := &model.Node{
			Name:  d.Name(),
			Path:  path,
			IsDir: d.IsDir(),
		}

		parentPath := filepath.Dir(path)
		if parentNode, ok := nodesByPath[parentPath]; ok {
			parentNode.Children = append(parentNode.Children, node)
		} else {
			return fmt.Errorf("internal walker error: could not find parent for path %s", path)
		}

		if d.IsDir() {
			nodesByPath[path] = node
		}

		return nil
	})

	if walkErr != nil {
		return nil, fmt.Errorf("error walking directory '%s': %w", rootPath, walkErr)
	}

	recursiveSort(rootNode)
	return rootNode, nil
}

// recursiveSort sorts the children of a Node alphabetically, ensuring that
// directories are always listed before files at the same level.
func recursiveSort(node *model.Node) {
	if !node.IsDir || len(node.Children) == 0 {
		return
	}

	sort.Slice(node.Children, func(i, j int) bool {
		if node.Children[i].IsDir != node.Children[j].IsDir {
			return node.Children[i].IsDir
		}
		return node.Children[i].Name < node.Children[j].Name
	})

	for _, child := range node.Children {
		if child.IsDir {
			recursiveSort(child)
		}
	}
}
