# Contributing

Keep changes scoped to one SDK when possible and run that SDK's checks before opening a pull request.

- Go: `cd golang && go test ./...`
- Python: `cd python && python3 -m pip install -r requirements.txt && python3 cli.py --help`
- TypeScript: `cd typescript && npm ci && npm run typecheck && npm run build && npm test`

Do not commit generated `dist/`, dependency directories, local environment files, or API credentials.
