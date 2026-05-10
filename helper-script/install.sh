#!/bin/sh
set -e

# Config
BASE_URL="https://dl.yourdomain.com/mycli/latest" #TODO Change when ready to move to prod
BINARY_NAME="mycli"

# 1. Detect OS/Arch
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
esac

# 2. Formulate download name
REMOTE_FILE="${BINARY_NAME}-${OS}-${ARCH}"
if [ "$OS" = "mingw" ] || [ "$OS" = "msys" ]; then
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