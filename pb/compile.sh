#!/usr/bin/env sh

protoc contentmgrsvc.proto --go_out=plugins=grpc:.
