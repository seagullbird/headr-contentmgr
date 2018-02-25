#!/usr/bin/env sh

protoc repoctlsvc.proto --go_out=plugins=grpc:.
