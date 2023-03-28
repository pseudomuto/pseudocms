db: cockroach start-single-node --insecure --host localhost
api: go run ./cmd/api -addr localhost:8000
gateway: go run ./cmd/gateway -addr localhost:9000 -rpc-host localhost:8000
