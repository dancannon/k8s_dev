//go:generate protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/gogo/protobuf/protobuf --gogo_out=. hello/hello.proto

package proto
