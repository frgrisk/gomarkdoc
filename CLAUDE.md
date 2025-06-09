# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gomarkdoc is a Go documentation generator that creates markdown documentation from Go source code. It serves as both a command-line tool and a programmatic library for generating markdown documentation for Go packages.

## Common Development Commands

### Build
```bash
mage build                    # Builds the binary to ./bin/gomarkdoc
```

### Test
```bash
mage test                     # Runs all tests with coverage
mage coverage                 # Opens HTML coverage report in browser
```

### Lint
```bash
mage lint                     # Runs golangci-lint (v1.51.1)
```

### Documentation
```bash
mage doc                      # Generates documentation for the project
mage docVerify               # Verifies documentation is up-to-date
mage regenerateTestDocs      # Regenerates test documentation in all formats
```

### Code Generation
```bash
mage generate                # Runs go generate for template embedding
```

## Architecture Overview

### Core Components

1. **Root Package** - Main renderer and template engine
   - `renderer.go` - Processes templates and generates documentation
   - `templates.go` - Embedded template definitions

2. **`/lang`** - Go language parsing and AST analysis
   - Parses Go packages using standard library tools
   - Creates structured representations of packages, types, functions, etc.

3. **`/format`** - Output formatting for different markdown flavors
   - GitHub, Azure DevOps, and plain markdown formats
   - Interface-based design for extensibility

4. **`/cmd/gomarkdoc`** - CLI implementation
   - Uses Cobra for command parsing
   - Supports configuration via `.gomarkdoc.yml`

5. **`/templates`** - Go text templates for rendering documentation elements

### Key Design Patterns

- **Template-based rendering** using Go's text/template package
- **Interface-based formatting** allows pluggable output formats
- **Self-documenting** - the project uses itself to generate its documentation
- **Embedded resources** - templates are compiled into the binary

## Important Notes

- The project requires Go 1.13.x or later and Mage 1.9.x or later
- Always run `mage docVerify` before committing to ensure docs are up-to-date
- Test documentation is generated in multiple formats (GitHub, Azure DevOps, plain) for comparison
- Configuration can be provided via CLI flags or `.gomarkdoc.yml` file
- The project uses golangci-lint for code quality checks