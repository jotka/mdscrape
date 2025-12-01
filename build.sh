#!/bin/bash

set -e

APP_NAME="mdscrape"
VERSION="${VERSION:-dev}"
BUILD_DIR="build"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log() {
    echo -e "${GREEN}[BUILD]${NC} $1"
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
    exit 1
}

# Clean build directory
clean() {
    log "Cleaning build directory..."
    rm -rf "$BUILD_DIR"
    mkdir -p "$BUILD_DIR"
}

# Build for current platform
build_current() {
    log "Building for current platform..."
    go build -ldflags="-s -w -X main.Version=$VERSION" -o "$APP_NAME" .
    log "Built: $APP_NAME"
}

# Build for all platforms
build_all() {
    clean

    platforms=(
        "darwin/amd64"
        "darwin/arm64"
        "linux/amd64"
        "linux/arm64"
        "windows/amd64"
    )

    for platform in "${platforms[@]}"; do
        GOOS="${platform%/*}"
        GOARCH="${platform#*/}"
        output="$BUILD_DIR/${APP_NAME}-${GOOS}-${GOARCH}"

        if [ "$GOOS" = "windows" ]; then
            output="${output}.exe"
        fi

        log "Building for $GOOS/$GOARCH..."
        GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w -X main.Version=$VERSION" -o "$output" .

        if [ $? -eq 0 ]; then
            log "Built: $output"
        else
            error "Failed to build for $GOOS/$GOARCH"
        fi
    done

    log "All builds complete!"
    ls -lh "$BUILD_DIR"
}

# Install locally
install_local() {
    build_current
    log "Installing to /usr/local/bin..."
    sudo mv "$APP_NAME" /usr/local/bin/
    log "Installed: $(which $APP_NAME)"
}

# Run tests
test() {
    log "Running tests..."
    go test -v ./...
}

# Show help
help() {
    echo "Usage: ./build.sh [command]"
    echo ""
    echo "Commands:"
    echo "  build     Build for current platform (default)"
    echo "  all       Build for all platforms (darwin, linux, windows)"
    echo "  install   Build and install to /usr/local/bin"
    echo "  clean     Remove build artifacts"
    echo "  test      Run tests"
    echo "  help      Show this help"
    echo ""
    echo "Environment variables:"
    echo "  VERSION   Set version string (default: dev)"
}

# Main
case "${1:-build}" in
    build)
        build_current
        ;;
    all)
        build_all
        ;;
    install)
        install_local
        ;;
    clean)
        clean
        log "Clean complete"
        ;;
    test)
        test
        ;;
    help|--help|-h)
        help
        ;;
    *)
        error "Unknown command: $1"
        ;;
esac
