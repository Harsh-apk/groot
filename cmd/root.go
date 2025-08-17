package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Variables to hold the version info passed from main.go
var (
	version string
	commit  string
	date    string
)

// SetVersionInfo is called by main.go to pass in build-time variables.
func SetVersionInfo(v, c, d string) {
	version = v
	commit = c
	date = d
}

// versionCmd defines the new "version" command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of Groot.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Groot version %s\n", version)
		fmt.Printf("Commit: %s, Built on: %s\n", commit, date)
	},
}

var asciiArt = `
 ██████╗ ██████╗  ██████╗  ██████╗ ████████╗
██╔════╝ ██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝
██║  ███╗██████╔╝██║   ██║██║   ██║   ██║   
██║   ██║██╔══██╗██║   ██║██║   ██║   ██║   
╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝   ██║   
 ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝ 
`

// grootBrown defines the custom brown color for the ASCII art.
var grootBrown = color.New().AddRGB(160, 82, 45).SprintFunc()

// aboutCmd defines the "about" command with a more detailed and professional layout.
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Shows information about Groot.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(grootBrown(asciiArt))
		fmt.Println("\n  A Codebase Analyzer for Enhanced LLM Context")
		fmt.Printf("  %-12s %s\n", "Version:", version) // Now uses the dynamic version
		fmt.Printf("  %-12s %s (%s)\n", "Author:", "Harsh", color.CyanString("https://github.com/harsh-apk"))
		fmt.Println("\n  Groot scans your source code to create a detailed, LLM-friendly")
		fmt.Println("  overview, improving the quality of AI-assisted development tasks.")
	},
}

// contributeCmd defines the "contribute" command with a consistent style.
var contributeCmd = &cobra.Command{
	Use:   "contribute",
	Short: "Shows information on how to contribute to Groot.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(grootBrown(asciiArt))
		color.Yellow("\nContributing to Groot\n\n")
		fmt.Println("  Thank you for your interest in making Groot better!")
		fmt.Println("  The best way to contribute is by reporting issues or submitting")
		fmt.Println("  pull requests on our GitHub repository.")
		fmt.Printf("  %-12s %s\n", "GitHub:", color.CyanString("https://github.com/harsh-apk/groot"))
	},
}

var rootCmd = &cobra.Command{
	Use:   "groot",
	Short: "Generates a comprehensive codebase overview for LLM prompting.",
	Long: `
A tool for generating a comprehensive overview of your codebase,
designed to enhance the context provided to Large Language Models (LLMs).

Run 'groot' to start an interactive session, or use 'groot [command]' for more options.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintIntro()
		// After the intro, the interactive session starts automatically.
		analyzeCmd.Run(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	// This line disables Cobra's default 'completion' command for a cleaner interface.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// Add the new versionCmd to the list of commands.
	rootCmd.AddCommand(analyzeCmd, aboutCmd, contributeCmd, versionCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// PrintIntro displays the ASCII art, welcome message, and command info.
func PrintIntro() {

	fmt.Println(grootBrown(asciiArt))
	color.Green("I am Groot! A tool to Analyze Codebase for Enhanced LLM Context\n\n")

	// Display available commands in a clean, consistent format.
	color.Yellow("Usage:")
	fmt.Println("  Run 'groot' to start the interactive analysis session.")
	fmt.Println("  Or use one of the following commands:")
	fmt.Printf("  %-12s  %s\n", color.CyanString("about"), "Learn more about the Groot tool.")
	fmt.Printf("  %-12s  %s\n", color.CyanString("contribute"), "Find out how to contribute.")
	fmt.Printf("  %-12s  %s\n", color.CyanString("version"), "Show the application version.")
	fmt.Printf("  %-12s  %s\n\n", color.CyanString("help"), "Show this help message.")

	color.HiBlack("Starting interactive session now...\n")
}
