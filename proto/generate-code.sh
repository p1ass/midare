#!/usr/bin/env sh
set -eu

echo "Installing dependencies..."

go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

echo "Generating twirp server code..."

PROTO_SRC_PATH="./"
PROTO_FILES="./*.proto"

# NOTE: Goのimportパスを正しくするために、オプションを設定している。
# https://twitchtv.github.io/twirp/docs/command_line.html#modifying-imports
GO_IMPORT_PREFIX="github.com/p1ass/midare"
GO_OUT_PATH="../backend"
protoc \
  --proto_path=$PROTO_SRC_PATH \
  --go_out=$GO_OUT_PATH \
  --go_opt=module=$GO_IMPORT_PREFIX \
  --twirp_out=module=$GO_IMPORT_PREFIX:$GO_OUT_PATH \
  $PROTO_FILES

echo "Generating documents..."

protoc \
  --proto_path=$PROTO_SRC_PATH \
  --doc_out=./docs \
  --doc_opt=markdown,index.md $PROTO_FILES