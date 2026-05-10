#!/bin/bash

# Configuration
APP_NAME="two-mi18n-helper-script"
DIST_DIR="./dist/latest"
VERSION="1.0.0"

# Clean previous builds
rm -rf ./dist
mkdir -p $DIST_DIR

# Platforms to build for
PLATFORMS=(
    "linux/amd64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

echo "Building $APP_NAME v$VERSION..."

for PLATFORM in "${PLATFORMS[@]}"; do
    OS="${PLATFORM%/*}"
    ARCH="${PLATFORM#*/}"
    OUTPUT_NAME="${APP_NAME}-${OS}-${ARCH}"
    
    if [ "$OS" = "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi

    echo "--- Building for $OS/$ARCH..."
    # -s -w removes debug info to shrink the binary by ~30%
    GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w" -o "$DIST_DIR/$OUTPUT_NAME" .
done

# Generate SHA256 Checksums
echo "Generating checksums..."
cd $DIST_DIR
sha256sum * > checksums.txt

# Copy install.sh to dist folder
cp ../../install.sh ./

echo "Build complete! Files are in $DIST_DIR"