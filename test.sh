#!/bin/bash

echo "🌙 Testing MoonWalk Enhanced Version"
echo "=================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first:"
    echo "   macOS: brew install go"
    echo "   Linux: sudo apt install golang-go"
    exit 1
fi

echo "✅ Go found: $(go version)"

# Build the project
echo "🔨 Building moonwalk..."
go mod tidy
go build -o moonwalk .

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

echo ""
echo "🧪 Running tests..."
echo ""

echo "1. Basic CLI usage:"
./moonwalk | head -10

echo ""
echo "2. With file sizes:"
./moonwalk -size | head -5

echo ""
echo "3. Summary statistics:"
./moonwalk -summary

echo ""
echo "4. Go files only:"
./moonwalk -ext .go

echo ""
echo "🎮 To test TUI mode, run:"
echo "   ./moonwalk -tui"
echo ""
echo "TUI Controls:"
echo "   ↑/↓ or j/k - Navigate"
echo "   / - Search files"
echo "   d - Toggle details"
echo "   r - Refresh"
echo "   q - Quit"