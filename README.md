# go-ai

A small, provider-neutral Go library for text and structured JSON generation with OpenAI Responses and Gemini generateContent. It requires Go 1.25 or later and has no external dependencies.

```sh
go get github.com/Dzarlax-AI/go-ai
```

Construct a provider with `openai.New(ai.Config{...})` or `gemini.New(ai.Config{...})`, then call `Generate` with a context and request. Each request specifies its model and prompt, and may specify temperature, output token limit, and a JSON Schema. See the [compilable text and structured JSON examples](examples/example_test.go); both use a synthetic HTTP transport and require no API key or network access.

`Config.APIKey` is required; the library does not load environment variables. `Config.HTTPClient` and `Config.BaseURL` are optional. An OpenAI BaseURL ends in `/v1` and a Gemini BaseURL ends in `/v1beta/models` (or another version's `/models`): supply the collection, not a per-model template. The default client has a 60-second timeout; callers may supply their own client and context deadlines.

Each call performs at most one HTTP attempt and does not follow redirects or retry. Callers own retry policy, application-level validation, and safe handling of generated text. Request and response bodies are bounded to 8 MiB and 2 MiB respectively. Errors do not include upstream bodies or transport URLs. A response may be complete, a refusal, or incomplete; refusal text is not exposed.

Schemas support object, array, string, number, integer and boolean, with properties, required, items, enum, description, nullable, boolean additionalProperties, and nested anyOf alternatives (including null). Combinations of anyOf with other validation keywords, $defs/$ref, and numeric or length limits are rejected. Types may be uppercase Gemini schema names. Unsupported constructs and schemas over 120,000 bytes are rejected before the HTTP request. OpenAI requires a root object, closed objects and every property required; originally optional properties become nullable. Callers must enforce their original constraints after generation. The library checks structured output for valid JSON, not conformance to the original schema.

Usage values are pointers: nil means the count is unknown. Cached input is included in input. OpenAI counts remain unknown when omitted. Gemini uses ProtoJSON, which can omit zero-valued scalar fields: when prompt tokens are reported, an omitted cached count means zero; when candidate tokens are reported, omitted thoughts mean zero. Gemini output is candidate tokens plus thoughts, with thoughts also reported as the reasoning subset. If the corresponding prompt or candidate count is absent, its derived subset and output remain unknown. Total is the provider-reported total, never a sum that duplicates cached or reasoning tokens. See [Gemini usage metadata](https://ai.google.dev/api/generate-content#UsageMetadata) and [ProtoJSON default values](https://protobuf.dev/programming-guides/json/#presence-and-default-values).

Licensed under the [MIT License](LICENSE).
