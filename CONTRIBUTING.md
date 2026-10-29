# Contributing

Use Go 1.25 or later. Changes should preserve the provider-neutral contract and avoid adding dependencies without a clear need. Tests must use synthetic responses and must not call real providers or require credentials.

Before submitting a change, run:

```sh
go test -race -count=1 ./...
go vet ./...
go mod verify
```

Document observable changes in the changelog and include tests for meaningful behavior changes.
