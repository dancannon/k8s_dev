#!/bin/sh -e
#
# Generate all etcd protobuf bindings.
# Run from repository root.
#
if ! protoc --version > /dev/null; then
    echo "could not find protoc, is it installed + in PATH?"
    exit 255
fi

if [[ "$unamestr" == 'Darwin' ]]; then
    find ${PROTO_DIR} -name '*.pb.go' | xargs rm -f
else
    find ${PROTO_DIR} -name '*.pb.go' | xargs -r rm -f
fi

find ${PROTO_DIR} -name '*.proto' -exec echo {} \;
find ${PROTO_DIR} -name '*.proto' -exec protoc -I$SRCPATH --go_out=${SRCPATH} {} \;
find ${PROTO_DIR} -name '*.proto' -exec protoc --gogo_out=. -I=.:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src **.*.proto
