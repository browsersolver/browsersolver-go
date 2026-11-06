# Examples

```sh
export BROWSERSOLVER_API_KEY="YOUR_API_KEY"
go run ./examples/json
```

- `json/` — call the featured endpoint and print the JSON body.
- `binary/` — save a binary response to a file.
- `request/` — call any path with `Request()` and handle API errors.
- `wait-for-job/` — start an async job and wait for its result.
