# Agent and contributor guidance

## Project contract

- Madmaid is a local-first Mermaid editor. Diagram name/source are stored in browser IndexedDB, not on the Go server.
- `src/` and `index.html` are source of truth. `dist/` is generated frontend output embedded by `//go:embed all:dist` in `main.go`.
- Go server listens on `127.0.0.1:8080` by default. Do not change network exposure defaults without explicit product/security justification.
- GitHub Pages deployment is static hosting only. Keep its Vite base path set to the repository subpath; executable builds require root-relative assets.

## Implementation expectations

- Preserve offline/local data behavior. Do not introduce server-side diagram storage or external runtime dependencies without explicit request.
- Keep Mermaid security level strict. Treat imported archives and diagram text as untrusted input; validate imported schema and avoid unsafe DOM insertion.
- Preserve autosave correctness when switching, deleting, restoring, exporting, or navigating between diagrams.
- Keep controls keyboard-accessible with names, titles/tooltips where appropriate, visible focus styles, and meaningful modal behavior.
- Do not edit generated files in `dist/` by hand. Do not add credentials, user backups, local browser data, or built executables to source control.
- Keep changes focused. Add/update tests when behavior changes and avoid weakening checks to obtain a green build.

## Local development

Requirements: Node.js 22.12+, npm, and Go 1.22+.

```sh
npm ci
npm run dev
```

## Required validation

For frontend or embedded asset changes:

```sh
npm ci
npm run build
go test ./...
go build -o /tmp/madmaid-check .
```

For changes to GitHub workflows, validate YAML syntax and verify workflow permissions, trigger branches/tags, artifact names, and deployment paths. For UI behavior changes, run browser E2E against the production build/embedded Go server when browser automation is available; report any environment limitation explicitly.

## CI and publishing

- `.github/workflows/ci.yml`: pull request and branch validation; least privilege permissions.
- `.github/workflows/release.yml`: publish native executable artifacts only from version tags.
- `.github/workflows/pages.yml`: publish static `dist/` with `/madmaid/` base to GitHub Pages. Deployment is opt-in through repository Pages settings.
- Avoid force pushes or destructive git operations. Before pushing, inspect branch, status, remotes, and diff; never overwrite existing user changes.
