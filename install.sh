#!/bin/bash
# DocxTerminal Installer — Linux/macOS/Git Bash

set -e

echo "📦 DocxTerminal Installer"
echo ""

# Check Go
if ! command -v go &> /dev/null; then
    echo "❌ Go not found. Install from https://go.dev/dl/ (1.24+)"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}')
echo "✓ Go $GO_VERSION found"

# Detect GOPATH/bin
GOBIN="${GOBIN:-$(go env GOPATH)/bin}"
mkdir -p "$GOBIN"
echo "✓ Install dir: $GOBIN"

# Check PATH
if [[ ":$PATH:" != *":$GOBIN:"* ]]; then
    echo ""
    echo "⚠️  $GOBIN not in PATH"
    echo ""
    echo "Add to ~/.bashrc or ~/.zshrc:"
    echo "  export PATH=\"\$GOBIN:\$PATH\""
    echo "  export GOBIN=\"\$GOBIN\""
    echo ""
    echo "Or run:"
    echo "  export PATH=\"$GOBIN:\$PATH\""
fi

# Build
echo ""
echo "🔨 Building dt..."
go build -o dt .

# Install
echo "📂 Installing to $GOBIN/dt..."
mv dt "$GOBIN/dt"
chmod +x "$GOBIN/dt"

# Verify
echo ""
if command -v dt &> /dev/null; then
    echo "✅ Install OK! Try: dt"
else
    echo "⚠️  dt not in PATH yet. Add $GOBIN to PATH, then restart terminal."
fi
