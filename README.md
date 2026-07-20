# ONR SDKs

Official multi-language SDKs and CLI clients for testing Open Next Router (ONR) OpenAI, Anthropic, and Gemini endpoints.

## Packages

- `golang/`: Go SDK and `onr-sdk` CLI.
- `python/`: Python CLI.
- `typescript/`: TypeScript SDK and `onr-sdk-ts` CLI.

Each package is independently installable from its own directory. See the package README for setup, environment variables, and examples.

## Development

```bash
git clone https://github.com/r9s-ai/onr-sdk.git
cd onr-sdk

(cd golang && go test ./...)
(cd python && python3 -m pip install -r requirements.txt && python3 cli.py --help)
(cd typescript && npm ci && npm run typecheck && npm test)
```

This repository does not publish packages yet. It provides source SDKs and CLI clients for ONR testing.
