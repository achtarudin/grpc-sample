install-tools:
	go install github.com/bufbuild/buf/cmd/buf@v1.59.0

clear-protogen:
	rm -rf ./protogen

generate-proto:
	buf generate --template buf.gen.yaml

run-cli:
	gow run ./cmd/cli/*.go