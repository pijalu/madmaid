# Madmaid

Madmaid is a self-hosted, local-first Mermaid diagram editor. Edit Mermaid source with a live preview, keep multiple diagrams in browser storage, and export or restore workspace backups. The Go server serves static assets only; diagram content stays in the browser.

## Features

- Live Mermaid preview with strict Mermaid security settings.
- Multiple browser-local diagrams with editable names and automatic saving.
- SVG and PNG export; PNG dimensions are configurable.
- Focus view with zoom controls; collapsible diagram list and source editor.
- JSON workspace backup and restore.
- Single Go executable with frontend assets embedded; no Node.js needed at runtime.
- Optional static hosting through GitHub Pages. Pages stores diagrams in each visitor's browser, not on GitHub.

## Requirements

- Node.js 22.12 or newer and npm (Mermaid 11 toolchain).
- Go 1.22 or newer for the server executable.

## Develop locally

```sh
npm ci
npm run dev
```

Open the local URL printed by Vite. The development server hosts the frontend only. Browser storage is origin-specific: changing hostname, port, browser, or protocol creates a separate workspace.

## Build and run

```sh
npm ci
npm run build
go build -o madmaid .
./madmaid
```

Visit <http://127.0.0.1:8080>. The server listens on loopback by default. To intentionally expose it to a network, use `./madmaid -addr 0.0.0.0:8080` and configure firewall access accordingly. The executable embeds `dist/`; build the frontend before `go build` whenever frontend files change.

Convenience targets:

```sh
make install  # npm install
make dev      # Vite development server
make build    # frontend and Go executable
make start    # build, then run executable
```

### Cross-compilation

Build native executables using Go's target variables. Build the frontend once, then for example:

```sh
GOOS=linux GOARCH=amd64 go build -o madmaid-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -o madmaid-linux-arm64 .
GOOS=darwin GOARCH=amd64 go build -o madmaid-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -o madmaid-darwin-arm64 .
GOOS=windows GOARCH=amd64 go build -o madmaid-windows-amd64.exe .
GOOS=windows GOARCH=arm64 go build -o madmaid-windows-arm64.exe .
```

GitHub Actions CI builds Linux, macOS, and Windows binaries for amd64 and arm64 and uploads them as workflow artifacts. See [Releases and CI](#releases-and-ci).

## Data, privacy, backups

- Diagrams, names, and source are stored in IndexedDB in the current browser and origin.
- The Go server does not receive diagram contents. GitHub Pages hosts only static frontend assets; data stays in each visitor's browser.
- Browser data can be cleared or evicted. Use **Download backup** regularly to save a JSON archive outside the browser.
- **Import archive** replaces the current browser's diagrams after confirmation. Keep a separate backup before restoring.
- SVG and PNG exports are images of a rendered diagram, not workspace backups.
- Do not place sensitive information in public issue reports or exported images/backups.

## GitHub Pages hosting

Pages can host the frontend as a static website. The deployment workflow publishes `dist/` from the `main` branch to GitHub Pages. It builds with the repository subpath (`/madmaid/`), which is needed for a project Pages URL. The same workflow can be run manually from Actions.

One-time repository setup:

1. Push the repository to GitHub.
2. In **Settings → Pages → Build and deployment**, select **GitHub Actions** as the source.
3. Ensure the default deployment branch is `main` (or update workflow branch filters if different).
4. Push to `main`, or manually run **Deploy GitHub Pages** in Actions.
5. Find the published URL under **Settings → Pages**; it is typically `https://pijalu.github.io/madmaid/`.

Pages deployment and native executable builds are separate jobs. Pages runs the Vite build with a `/madmaid/` base path; executable jobs build Vite assets at `/` before embedding them. A static Pages deployment does not run the Go server.

## Releases and CI

`.github/workflows/ci.yml` runs frontend production build, Go tests, and Go build on pushes and pull requests. `.github/workflows/release.yml` builds Linux/macOS/Windows amd64 and arm64 executables for version tags matching `v*`, then attaches binaries and checksums to a GitHub Release. The build matrix runs frontend build on each runner before compiling the embedded executable.

Create a release by pushing a version tag, for example:

```sh
git tag v1.0.0
git push origin v1.0.0
```

## Development and project layout

```text
main.go          Go static-file server; embeds dist/
src/main.ts      UI, Mermaid rendering, IndexedDB, import/export
src/style.css    UI styles
index.html       Vite entry document
Makefile         Local build/run shortcuts
.github/         CI, release, and Pages workflows
```

Source files are authoritative. `dist/` is generated output and should not be edited directly. The CI build checks TypeScript and creates optimized Vite assets; Vite may report a large-bundle advisory due to Mermaid's bundled diagram implementations.

Before submitting changes, run:

```sh
npm ci
npm run build
go test ./...
go build -o /tmp/madmaid-check .
```

See [AGENTS.md](AGENTS.md) for repository-specific implementation and validation guidance.
