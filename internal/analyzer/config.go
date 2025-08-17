package analyzer

import "github.com/harsh-apk/groot/internal/model"

// CompiledLanguageConfig holds the static language configuration, removing the need for an external languages.yml file.
// This configuration is compiled directly into the binary.
var CompiledLanguageConfig = model.LanguageConfig{
	Languages: []model.Language{
		{
			Name:           "Go",
			FileExtensions: []string{".go"},
			Queries: []model.LanguageQuery{
				{Type: "Function", Query: `(function_declaration name: (identifier) @name)`},
				{Type: "Method", Query: `(method_declaration name: (field_identifier) @name)`},
				{Type: "Interface", Query: `(type_spec name: (type_identifier) @name (interface_type))`},
				{Type: "Struct", Query: `(type_spec name: (type_identifier) @name (struct_type))`},
			},
		},
		{
			Name:           "JavaScript",
			FileExtensions: []string{".js", ".jsx", ".mjs", ".cjs"},
			Queries: []model.LanguageQuery{
				{Type: "Component", Query: `(export_statement declaration: (lexical_declaration (variable_declarator name: (identifier) @name value: (arrow_function))))`},
				{Type: "Component", Query: `(lexical_declaration (variable_declarator name: (identifier) @name value: (arrow_function)))`},
				{Type: "Component", Query: `(export_statement declaration: (function_declaration name: (identifier) @name))`},
				{Type: "Constant", Query: `(export_statement declaration: (lexical_declaration (variable_declarator name: (identifier) @name)))`},
				{Type: "Function", Query: `(function_declaration name: (identifier) @name)`},
				{Type: "Component", Query: `(export_statement value: (identifier) @name)`},
				{Type: "Class", Query: `(class_declaration name: (identifier) @name)`},
				{Type: "Class Component", Query: `(export_statement declaration: (class_declaration name: (identifier) @name))`},
				{Type: "Method", Query: `(method_definition name: (property_identifier) @name)`},
			},
		},
		{
			Name:           "Java",
			FileExtensions: []string{".java"},
			Queries: []model.LanguageQuery{
				{Type: "Controller", Query: `((class_declaration (modifiers (annotation name: (identifier) @ann)) name: (identifier) @name) (#eq? @ann "RestController"))`},
				{Type: "Service", Query: `((class_declaration (modifiers (annotation name: (identifier) @ann)) name: (identifier) @name) (#eq? @ann "Service"))`},
				{Type: "Repository", Query: `((class_declaration (modifiers (annotation name: (identifier) @ann)) name: (identifier) @name) (#eq? @ann "Repository"))`},
				{Type: "Class", Query: `(class_declaration name: (identifier) @name)`},
				{Type: "Method", Query: `(method_declaration name: (identifier) @name)`},
				{Type: "Interface", Query: `(interface_declaration name: (identifier) @name)`},
			},
		},
		{
			Name:           "Python",
			FileExtensions: []string{".py"},
			Queries: []model.LanguageQuery{
				{Type: "Function", Query: `(function_definition name: (identifier) @name)`},
				{Type: "Class", Query: `(class_definition name: (identifier) @name)`},
			},
		},
		{
			Name:           "Rust",
			FileExtensions: []string{".rs"},
			Queries: []model.LanguageQuery{
				{Type: "Function", Query: `(function_item name: (identifier) @name)`},
				{Type: "Struct", Query: `(struct_item name: (type_identifier) @name)`},
				{Type: "Enum", Query: `(enum_item name: (type_identifier) @name)`},
				{Type: "Trait", Query: `(trait_item name: (type_identifier) @name)`},
			},
		},
		{
			Name:           "HTML",
			FileExtensions: []string{".html", ".htm"},
			Queries:        []model.LanguageQuery{},
		},
		{
			Name:           "CSS",
			FileExtensions: []string{".css"},
			Queries: []model.LanguageQuery{
				{Type: "Class Selector", Query: `(class_selector) @name`},
				{Type: "ID Selector", Query: `(id_selector) @name`},
			},
		},
	},
}
