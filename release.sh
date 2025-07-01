#!/bin/bash

# Release script for DocumentFactory MCP Server
# This script automates the release process

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

# Check if version argument is provided
if [ $# -eq 0 ]; then
    print_error "Usage: $0 <version>"
    print_error "Example: $0 1.0.0"
    exit 1
fi

VERSION=$1
TAG="v$VERSION"

# Validate version format (basic check)
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    print_error "Invalid version format. Use semantic versioning (e.g., 1.0.0)"
    exit 1
fi

echo "DocumentFactory MCP Server Release Process"
echo "=========================================="
echo "Version: $VERSION"
echo "Tag: $TAG"
echo ""

# Check if we're on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    print_warning "Not on main branch (current: $CURRENT_BRANCH)"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Check for uncommitted changes
if [ -n "$(git status --porcelain)" ]; then
    print_error "You have uncommitted changes. Please commit or stash them first."
    git status --short
    exit 1
fi

print_status "Working directory is clean"

# Check if tag already exists
if git rev-parse "$TAG" >/dev/null 2>&1; then
    print_error "Tag $TAG already exists"
    exit 1
fi

print_status "Tag $TAG is available"

# Run tests
echo "Running tests..."
if ! make test-all; then
    print_error "Tests failed"
    exit 1
fi

print_status "All tests passed"

# Update NPM package version
echo "Updating NPM package version..."
cd npm-package
npm version $VERSION --no-git-tag-version
cd ..

print_status "NPM package version updated"

# Commit version changes
git add npm-package/package.json
git commit -m "Bump version to $VERSION"

print_status "Version changes committed"

# Create and push tag
git tag $TAG
git push origin main
git push origin $TAG

print_status "Tag $TAG created and pushed"

echo ""
echo "Release process completed!"
echo ""
echo "Next steps:"
echo "1. GitHub Actions will automatically build and release the Go binaries"
echo "2. GitHub Actions will automatically publish the NPM package"
echo "3. Monitor the Actions tab for deployment status"
echo ""
echo "Release URL: https://github.com/GlobalCents-DocumentFactory/mcp-documentfactory/releases/tag/$TAG"
