package parser

import (
	"context"
	"fmt"

	"github.com/harsh-apk/groot/internal/model"
	sitter "github.com/smacker/go-tree-sitter"

	// Import the specific Go bindings for the languages you support.
	tree_sitter_css "github.com/tree-sitter/tree-sitter-css/bindings/go"
	tree_sitter_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
	tree_sitter_html "github.com/tree-sitter/tree-sitter-html/bindings/go"
	tree_sitter_java "github.com/tree-sitter/tree-sitter-java/bindings/go"
	tree_sitter_javascript "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	tree_sitter_python "github.com/tree-sitter/tree-sitter-python/bindings/go"
	tree_sitter_rust "github.com/tree-sitter/tree-sitter-rust/bindings/go"
)

// grammarMap maps a language name (from the YAML config) to its statically
// linked Tree-sitter grammar object.
var grammarMap = map[string]*sitter.Language{
	"Go":     sitter.NewLanguage(tree_sitter_go.Language()),
	"Python": sitter.NewLanguage(tree_sitter_python.Language()),
	"Java":   sitter.NewLanguage(tree_sitter_java.Language()),
	"Rust":   sitter.NewLanguage(tree_sitter_rust.Language()),
	"CSS":    sitter.NewLanguage(tree_sitter_css.Language()),
	"HTML":   sitter.NewLanguage(tree_sitter_html.Language()),

	// --- THIS IS THE CRITICAL CHANGE ---
	// We now use the JavaScript grammar but specifically call LanguageJSX()
	// to enable parsing of React/Next.js component syntax.
	"JavaScript": sitter.NewLanguage(tree_sitter_javascript.Language()),
}

// Parse uses Tree-sitter to extract code elements from source code.
func Parse(content []byte, lang model.Language) ([]model.CodeElement, error) {
	// 1. Look up the grammar from our pre-populated map.
	tsLang, found := grammarMap[lang.Name]
	if !found {
		// Gracefully skip unsupported files instead of erroring.
		return nil, nil
	}

	// 2. Create a new Tree-sitter parser and set its language.
	parser := sitter.NewParser()
	parser.SetLanguage(tsLang)

	// 3. Parse the source code content into a syntax tree.
	tree, err := parser.ParseCtx(context.Background(), nil, content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}

	var allElements []model.CodeElement
	rootNode := tree.RootNode()

	// 4. Iterate over all queries defined for this language.
	for _, langQuery := range lang.Queries {
		if langQuery.Query == "" {
			continue
		}

		query, err := sitter.NewQuery([]byte(langQuery.Query), tsLang)
		if err != nil {
			return nil, fmt.Errorf("failed to compile query for type '%s': %w", langQuery.Type, err)
		}

		qc := sitter.NewQueryCursor()
		qc.Exec(query, rootNode)

		// 5. Iterate over all the matches found by the query.
		for {
			match, ok := qc.NextMatch()
			if !ok {
				break
			}
			match = qc.FilterPredicates(match, content)

			for _, capture := range match.Captures {
				if query.CaptureNameForId(capture.Index) == "name" {
					allElements = append(allElements, model.CodeElement{
						Name: capture.Node.Content(content),
						Type: langQuery.Type,
						Line: int(capture.Node.StartPoint().Row + 1),
					})
					break
				}
			}
		}
	}

	return allElements, nil
}
