#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# --- Step 1: Setup ---
echo "▶️ Step 1: Setting up the build directory..."
mkdir -p builds
echo "✅ Done."
echo

# --- Step 2: Install Cross-Compiler Toolchains ---
echo "▶️ Step 2: Installing cross-compilers via Homebrew..."

# For Linux: Use the messense tap
brew tap messense/macos-cross-toolchains
brew install x86_64-unknown-linux-gnu
brew install aarch64-unknown-linux-gnu

# For Windows (amd64 only): Use the standard mingw-w64 from Homebrew core
brew install mingw-w64

echo "✅ All toolchains are ready."
echo

# --- Step 3: Compile Binaries ---
echo "▶️ Step 3: Compiling binaries for all supported platforms..."

# Add ldflags to strip debug info and reduce binary size
LDFLAGS="-s -w"

# Build for Linux
echo "  -> Building for Linux (amd64)..."
CC=x86_64-unknown-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o builds/groot-linux-amd64 .

echo "  -> Building for Linux (arm64)..."
CC=aarch64-unknown-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="$LDFLAGS" -o builds/groot-linux-arm64 .

# Build for Windows
echo "  -> Building for Windows (amd64)..."
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags="$LDFLAGS" -o builds/groot-windows-amd64.exe .

# Build for macOS (no cross-compiler needed as we are on macOS)
echo "  -> Building for macOS (amd64 - Intel)..."
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="$LDFLAGS" -o builds/groot-darwin-amd64 .

echo "  -> Building for macOS (arm64 - Apple Silicon)..."
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -ldflags="$LDFLAGS" -o builds/groot-darwin-arm64 .

echo "✅ All binaries compiled successfully."
echo

# --- Step 4: Create Distributable Archives ---
echo "▶️ Step 4: Packaging binaries into archives..."

cd builds

# Create tarballs for Linux and macOS
tar -czvf groot-linux-amd64.tar.gz groot-linux-amd64
tar -czvf groot-linux-arm64.tar.gz groot-linux-arm64
tar -czvf groot-darwin-amd64.tar.gz groot-darwin-amd64
tar -czvf groot-darwin-arm64.tar.gz groot-darwin-arm64

# Create zip file for Windows
zip groot-windows-amd64.zip groot-windows-amd64.exe

cd ..
echo "✅ Archives created."
echo

# --- All Done ---
echo "🎉 Success! All binaries are built and archived in the 'builds' directory."