CRUSH.md — Quickstart for agents working in this repo

Build/Test/Lint
- Build (local): task build (uses go build -o injest cmd/injest/main.go)
- Clean: task clean
- Tidy modules: task tidy (runs go mod tidy)
- Test: no tests yet; scaffold with go test ./... when added
- Run a single test (when tests exist): go test ./... -run ^TestName$
- Lint (CI): golangci-lint runs via .github/workflows/golangci-lint.yml; locally use golangci-lint run (install if missing)
- Security/CodeQL: .github/workflows/codeql-analysis.yml
- Release build: goreleaser release --clean (see .goreleaser.yml) or goreleaser build --snapshot

Language/Code Style (Go 1.21)
- Formatting: go fmt ./... and goimports conventions; keep imports grouped stdlib, third-party, module
- Modules: defined in go.mod (module github.com/conacademy/injest); keep dependencies minimal, run task tidy after changes
- Types and naming: exported identifiers use PascalCase with doc comments; unexported use lowerCamelCase; avoid stutter; interfaces end with -er when natural (e.g., Jester)
- Errors: prefer returning errors; for CLI main(), write errors to stderr and exit non-zero; no panics in library code
- Logging/IO: CLI writes primary output to stdout, diagnostics to stderr; use fmt.Fprintf with os.Stdout/os.Stderr as in cmd/injest/main.go
- Strings: use strings.Builder for concatenation in loops or multi-part assembly (pkg/jester follows this)
- Flags: use spf13/pflag (already in use); keep short and long variants, provide --help, default-safe behavior
- Package layout: cmd/injest for CLI entrypoint; library code lives under pkg/ (pkg/jester). Add new reusable code under pkg/<name>

Testing Guidelines
- Use standard testing package; place *_test.go next to code; table-driven tests preferred
- For CLI behavior, test via small library functions under pkg where possible; keep main minimal and thin

Editor/Assistant Rules
- Cursor/Copilot: no project-specific rules found (.cursor/rules, .cursorrules, .github/copilot-instructions.md absent)
- Follow Go idioms and keep files comment-light; do not commit secrets; respect CODE_OF_CONDUCT.md tone

Handy commands
- Run: ./injest -h after build, or go run ./cmd/injest -h
- List jesters: ./injest --list
