#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' . && \
docker build -t 10.0.2.2:5000/hello-service . && \
docker push 10.0.2.2:5000/hello-service
