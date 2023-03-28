db: cockroach start-single-node --insecure --listen-addr localhost:26257
api: go run ./cmd/api -addr localhost:8000
gateway: go run ./cmd/gateway -addr localhost:9000 -rpc-host localhost:8000
docs: cd openapi && python -m http.server --bind localhost 9001
