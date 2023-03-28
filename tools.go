//go:build tools

package pseudocms

import (
	_ "github.com/gobuffalo/pop/v6/soda"                                  // for soda CLI
	_ "github.com/golang/mock/mockgen"                                    // for mockgen
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway" // for protoc-gen-grpc-gateway
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"    // for protoc-gen-openapiv2
	_ "github.com/mattn/goreman"                                          // for goreman
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"                     // for protoc-gen-go-grpc
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"                      // for protoc-gen-go
)
