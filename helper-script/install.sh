#!/bin/sh
set -e

# Config
BASE_URL="http://two-mi18n-helper-script.nicolasrenault.com/"
BINARY_NAME="two-mi18n-helper-script"

# Detect OS
OS_RAW="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS_RAW" in
    linux*)   OS="linux" ;;
    darwin*)  OS="darwin" ;;
    mingw*|msys*|cygwin*) OS="windows" ;; # Standardizes all Windows environments
    *) echo "Unsupported OS: $OS_RAW" && exit 1 ;;
esac

# Detect Architecture
ARCH_RAW="$(uname -m)"
case "$ARCH_RAW" in
    x86_64|amd64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH_RAW" && exit 1 ;;
esac

# Now REMOTE_FILE will correctly be "two-mi18n-helper-script-windows-amd64.exe"
REMOTE_FILE="${BINARY_NAME}-${OS}-${ARCH}"

if [ "$OS" = "windows" ]; then
    REMOTE_FILE="${REMOTE_FILE}.exe"
fi

# 3. Download to temp
echo "Downloading $REMOTE_FILE..."
curl -fsSL "$BASE_URL/$REMOTE_FILE" -o "/tmp/$BINARY_NAME"
chmod +x "/tmp/$BINARY_NAME"

# 4. Move to bin (requires sudo if not writable)
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"
if [ -w "/usr/local/bin" ]; then
    mv "/tmp/$BINARY_NAME" "$INSTALL_PATH"
else
    sudo mv "/tmp/$BINARY_NAME" "$INSTALL_PATH"
fi

echo "Installed successfully! Try running: $BINARY_NAME --help"