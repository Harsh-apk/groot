package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/harsh-apk/groot/internal/analyzer"
	"github.com/harsh-apk/groot/internal/model"
	"github.com/spf13/cobra"
)

// This struct will hold the answers from the interactive survey.
type analysisAnswers struct {
	Path            string
	SkipDirs        string
	IncludeExts     string
	Format          string
	OutputDirectory string
	OutputFileName  string
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyzes a codebase through an interactive session.",
	Run: func(cmd *cobra.Command, args []string) {
		answers, err := askAnalysisQuestions()
		if err != nil {
			// This can happen if the user cancels (e.g., Ctrl+C).
			fmt.Println("\nAnalysis cancelled.")
			return
		}

		// Process the comma-separated strings into slices.
		skipList := processStringList(answers.SkipDirs)
		includeList := processStringList(answers.IncludeExts)

		fmt.Println("\n🔍 Starting analysis...")
		rootNode, stats, err := analyzer.Analyze(answers.Path, skipList, includeList)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during analysis: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Analysis complete!")

		var finalOutput []byte

		// Format the output based on the user's choice.
		if answers.Format == "json" {
			result := model.AnalysisResult{Root: rootNode, Analytics: stats}
			finalOutput, _ = json.MarshalIndent(result, "", "  ")
		} else {
			treeOutput, analyticsOutput := analyzer.FormatText(rootNode, stats, includeList)
			finalOutput = []byte(treeOutput + analyticsOutput + time.Now().Format("\n\nLast Analysis completed at: 2006-01-02 15:04:05"))
		}

		// --- UPDATED: Write to file or print to console ---
		if answers.OutputFileName != "" {
			// Automatically add the correct file extension.
			fullFileName := fmt.Sprintf("%s.%s", answers.OutputFileName, answers.Format)
			fullPath := filepath.Join(answers.OutputDirectory, fullFileName)

			// Ensure the output directory exists.
			if err := os.MkdirAll(answers.OutputDirectory, os.ModePerm); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating output directory %s: %v\n", answers.OutputDirectory, err)
				os.Exit(1)
			}

			// Write the file.
			err := os.WriteFile(fullPath, finalOutput, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", fullPath, err)
				os.Exit(1)
			}
			fmt.Printf("\nOutput successfully written to %s\n", fullPath)
		} else {
			fmt.Println(string(finalOutput))
		}
	},
}

// askAnalysisQuestions defines and runs the interactive survey.
func askAnalysisQuestions() (*analysisAnswers, error) {
	// Get the current working directory as the default path.
	currentDir, _ := os.Getwd()

	questions := []*survey.Question{
		{
			Name: "path",
			Prompt: &survey.Input{
				Message: "Enter the path to the directory you want to analyze \n(press enter if you want to analyze the current directory) \n",
				Default: currentDir,
				Help:    "This is the root directory of the codebase you want to scan.",
			},
			Validate: survey.Required,
		},
		{
			Name: "skipDirs",
			Prompt: &survey.Input{
				Message: "Directories to skip (comma-separated, press Enter to skip):",
				Help:    "List any directories you want to exclude from the analysis, like 'node_modules' or 'dist'.",
			},
		},
		{
			Name: "includeExts",
			Prompt: &survey.Input{
				Message: "File extensions to include (e.g., .go,.js, press Enter for all):",
				Help:    "Specify which file types to focus on. If blank, all supported files will be analyzed.",
			},
		},
		{
			Name: "format",
			Prompt: &survey.Select{
				Message: "Choose an output format:",
				Options: []string{"txt", "json"},
				Default: "txt",
				Help:    "Choose 'txt' for a human-readable report or 'json' for machine-readable output.",
			},
		},
		{
			Name: "outputDirectory",
			Prompt: &survey.Input{
				Message: "Enter an output directory (press Enter to use current directory):",
				Default: currentDir,
				Help:    "The directory where the output file will be saved.",
			},
		},
		{
			Name: "outputFileName",
			Prompt: &survey.Input{
				Message: "Enter a base file name (press Enter to print to console):",
				Help:    "The result will be saved here. The correct extension (.txt or .json) will be added automatically.",
			},
		},
	}

	answers := &analysisAnswers{}
	err := survey.Ask(questions, answers)
	return answers, err
}

// processStringList is a helper function to convert a comma-separated string to a slice.
func processStringList(input string) []string {
	if input == "" {
		return []string{}
	}
	parts := strings.Split(input, ",")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}
