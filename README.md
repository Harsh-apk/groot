<div align="center">

![](/readmeImages/groot.png)

# Groot
**An intelligent codebase analyzer for enhanced LLM context.**

<div align="center">
<img src="images/groot.png"  height="400" />

</div>

<!-- <p>
  <a href="https://github.com/harsh-apk/groot/releases"><img src="https://img.shields.io/github/v/release/harsh-apk/groot" alt="Release"></a>
  <a href="https://github.com/harsh-apk/groot/actions"><img src="https://github.com/harsh-apk/groot/actions/workflows/release.yml/badge.svg" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/harsh-apk/groot"><img src="https://goreportcard.com/badge/github.com/harsh-apk/groot" alt="Go Report Card"></a>
  <a href="https://github.com/harsh-apk/groot/blob/main/LICENSE"><img src="https://img.shields.io/github/license/harsh-apk/groot" alt="License"></a>
</p> -->
</div>

Groot is an interactive command-line tool that scans any codebase and generates a comprehensive, context-rich overview. It builds a detailed file tree and outlines key code elements (functions, classes, etc.), producing a single, clean output perfect for pasting into an LLM like Gemini, Grok, Mistral, Deepseek, Claude, GPT etc.

This allows the LLM to understand your codebase's architecture, leading to better, more accurate responses for debugging, refactoring, and feature implementation.

### Groot in Action
<div align="center">

![Groot in Action](/images/ss.png)

</div>

---

### ✨ Features

* **Guided Interactive Experience:** A friendly CLI that walks you through the analysis process.
* **Broad Language Support:** Analyzes Go, Python, JavaScript (JSX/TSX), Java, Rust, and more out-of-the-box.
* **Smart & Customizable:** Intelligently ignores irrelevant files (`.git`, `node_modules`) and lets you customize the scan.
* **Multiple Formats:** Outputs to a clean text format for LLMs or JSON for tool integration.
* **Codebase Analytics:** Provides a quick summary of file counts, lines of code, and identified code elements.

### 🚀 Installation

#### macOS (Homebrew)
```sh
brew tap harsh-apk/groot
brew install groot
```

#### Windows (Scoop)
```sh
scoop bucket add groot [https://github.com/harsh-apk/scoop-groot.git](https://github.com/harsh-apk/scoop-groot.git)
scoop install groot
```


### 💡 Usage

Simply run `groot` in your project's root directory to start the interactive session.
```sh
groot
```
The tool will guide you with a series of questions to configure the analysis.

**Other Commands:**
* `groot about`: Shows information about the tool.
* `groot version`: Prints the current version.
* `groot contribute`: Provides information on how to contribute.

### 📄 Example Output

The generated text output is clean, simple, and ready to be used as LLM context.

```text
Codebase overview for: /Users/harsh/Desktop/groot

/Users/harsh/Desktop/groot
├── builds
│   ├── groot-darwin-amd64
│   ├── groot-darwin-amd64.tar.gz
│   ├── groot-darwin-arm64
│   ├── groot-darwin-arm64.tar.gz
│   ├── groot-linux-amd64
│   ├── groot-linux-amd64.tar.gz
│   ├── groot-linux-arm64
│   ├── groot-linux-arm64.tar.gz
│   ├── groot-windows-amd64.exe
│   └── groot-windows-amd64.zip
├── cmd
│   ├── analyze.go
│   │     - Struct: analysisAnswers (L18)
│   │     - Function: askAnalysisQuestions (L87)
│   │     - Function: processStringList (L147)
│   └── root.go
│         - Function: SetVersionInfo (L19)
│         - Function: Execute (L91)
│         - Function: PrintIntro (L103)
├── config
│   └── languages.yml
├── images
│   └── groot.png
├── internal
│   ├── analyzer
│   │   ├── analyzer.go
│   │   │     - Function: GetLanguageByFileExtension (L20)
│   │   │     - Function: Analyze (L36)
│   │   │     - Function: FormatText (L84)
│   │   │     - Function: worker (L97)
│   │   │     - Function: aggregateAnalytics (L120)
│   │   │     - Function: appendAnalytics (L149)
│   │   │     - Function: collectFileNodes (L187)
│   │   │     - Function: formatTree (L201)
│   │   │     - Function: directoryContainsIncludedFiles (L261)
│   │   └── config.go
│   ├── model
│   │   └── models.go
│   │         - Struct: LanguageQuery (L6)
│   │         - Struct: Language (L12)
│   │         - Struct: LanguageConfig (L19)
│   │         - Struct: CodeElement (L24)
│   │         - Struct: Node (L31)
│   │         - Struct: LanguageStats (L41)
│   │         - Struct: Analytics (L48)
│   │         - Struct: AnalysisResult (L59)
│   ├── parser
│   │   └── treesitter.go
│   │         - Function: Parse (L37)
│   └── walker
│       └── walker.go
│             - Function: BuildFileTree (L62)
│             - Function: recursiveSort (L138)
├── LICENSE
├── README.md
├── build.sh
├── go.mod
├── go.sum
├── groot
├── main.go
│     - Function: main (L14)
└── out.txt


---

📊 Analysis Report

Overall Summary
────────────────────────────────────────
Analysis Duration:   15ms
Files Scanned:       27
Files Parsed:        8
Total Lines of Code: 971
Total Elements Found: 27
────────────────────────────────────────

Language Breakdown
────────────────────────────────────────
▶ Go (8 files, 971 LOC)
  - Function:          18
  - Struct:            9



Last Analysis completed at: 2025-08-17 20:02:24
```

### 🤝 Contributing

Contributions are welcome! Whether it's reporting a bug, suggesting a feature, or submitting a pull request, all contributions are appreciated. Please see the `groot contribute` command for more details.

### 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.