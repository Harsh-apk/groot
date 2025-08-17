package model

import "time"

// LanguageQuery defines a specific Tree-sitter query.
type LanguageQuery struct {
	Type  string `json:"type"`
	Query string `json:"query"`
}

// Language represents the configuration for a programming language.
type Language struct {
	Name           string          `json:"name"`
	FileExtensions []string        `json:"file_extensions"`
	Queries        []LanguageQuery `json:"queries,omitempty"`
}

// LanguageConfig holds all language configurations.
type LanguageConfig struct {
	Languages []Language `json:"languages"`
}

// CodeElement represents a single parsed entity from a source code file.
type CodeElement struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Line int    `json:"line"`
}

// Node represents a single item in the file system tree.
type Node struct {
	Name         string        `json:"name"`
	Path         string        `json:"path"`
	IsDir        bool          `json:"is_dir"`
	LOC          int           `json:"lines_of_code,omitempty"`
	Children     []*Node       `json:"children,omitempty"`
	CodeElements []CodeElement `json:"elements,omitempty"`
}

// LanguageStats holds analytics for a specific language.
type LanguageStats struct {
	FileCount     int            `json:"file_count"`
	LOC           int            `json:"lines_of_code"`
	ElementCounts map[string]int `json:"element_counts,omitempty"`
}

// Analytics holds comprehensive statistics about the analysis process.
type Analytics struct {
	FilesScanned     int                      `json:"files_scanned"`
	FilesParsed      int                      `json:"files_parsed"`
	TotalLOC         int                      `json:"total_lines_of_code"`
	TotalElements    int                      `json:"total_elements"`
	PerLanguageStats map[string]LanguageStats `json:"language_stats,omitempty"`
	Duration         time.Duration            `json:"duration_nanoseconds"`
	DurationReadable string                   `json:"duration_readable"`
}

// AnalysisResult is the top-level struct for JSON output.
type AnalysisResult struct {
	Root      *Node     `json:"tree"`
	Analytics Analytics `json:"analytics"`
}
