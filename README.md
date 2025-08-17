<div align="center">
<img src="https://www.google.com/search?q=https://i.imgur.com/your-groot-logo.png" alt="Groot Logo" width="150"/>
<h1>Groot</h1>
<p><strong>An intelligent codebase analyzer for enhanced LLM context.</strong></p>

<p>
<a href="https://www.google.com/search?q=https://github.com/harsh-apk/groot/releases"><img src="https://www.google.com/search?q=https://img.shields.io/github/v/release/harsh-apk/groot" alt="Release"></a>
<a href="https://www.google.com/search?q=https://github.com/harsh-apk/groot/actions"><img src="https://www.google.com/search?q=https://github.com/harsh-apk/groot/actions/workflows/release.yml/badge.svg" alt="Build Status"></a>
<a href="https://www.google.com/search?q=https://goreportcard.com/report/github.com/harsh-apk/groot"><img src="https://www.google.com/search?q=https://goreportcard.com/badge/github.com/harsh-apk/groot" alt="Go Report Card"></a>
<a href="https://www.google.com/search?q=https://github.com/harsh-apk/groot/blob/main/LICENSE"><img src="https://www.google.com/search?q=https://img.shields.io/github/license/harsh-apk/groot" alt="License"></a>
</p>
</div>
🌳 What is Groot?

Groot is an interactive command-line tool that scans any codebase and generates a comprehensive, context-rich overview. It creates a detailed file tree and outlines key code elements (functions, classes, components, etc.), producing a single, clean output perfect for pasting into a Large Language Model (LLM) like GPT-4, Claude, or Gemini.
The Problem

When working with LLMs, providing good context is everything. Pasting individual files is inefficient and loses the high-level structure of the project.
The Solution

Groot walks through your entire project, intelligently ignores irrelevant files, parses the code, and combines the structure and key elements into a single, concise output. This allows the LLM to understand your codebase's architecture, leading to better, more accurate responses for debugging, refactoring, and feature implementation.
✨ Features

    Interactive CLI: A user-friendly, question-and-answer interface that guides you through the analysis.

    Multi-Language Support: Out-of-the-box support for Go, JavaScript (including JSX), Java, Python, Rust, and more.

    Customizable Analysis: Interactively choose which directories to skip or which file extensions to include.

    Multiple Output Formats: Generate a human-readable text report or a machine-readable json output for tool integration.

    Detailed Analytics: Get a summary of your codebase, including lines of code, file counts, and a breakdown of code elements by language.

    Self-Contained: The language parsing rules are compiled directly into the binary, so there are no extra configuration files to manage.

🚀 Installation
macOS (Homebrew)

brew tap harsh-apk/homebrew-groot
brew install groot

Windows (Scoop)

scoop bucket add groot-bucket [https://github.com/harsh-apk/scoop-bucket.git](https://github.com/harsh-apk/scoop-bucket.git)
scoop install groot

Linux (APT and YUM)

Download the .deb or .rpm package from the latest release and install it with your system's package manager.
Docker

docker pull ghcr.io/harsh-apk/groot:latest
docker run -it --rm -v "$(pwd)":/app ghcr.io/harsh-apk/groot:latest

Manual Installation

Download the pre-compiled binary for your operating system from the Releases page, extract it, and place it in your PATH.
kullanım

Simply run groot in your terminal to start the interactive session:

groot

The tool will then ask you a series of questions to configure the analysis:

<div align="center">
<img src="https://www.google.com/search?q=https://i.imgur.com/your-interactive-demo.gif" alt="Groot Interactive Demo"/>
</div>
Commands

    groot: Starts the interactive analysis.

    groot about: Shows information about the tool.

    groot contribute: Provides information on how to contribute.

    groot version: Prints the current version.

💡 Example Output
Text Format

🤝 Contributing

Contributions are welcome! Whether it's reporting a bug, suggesting a feature, or submitting a pull request, all contributions are appreciated. Please see the groot contribute command for more details.
📜 License

This project is licensed under the MIT License. See the LICENSE file for details.