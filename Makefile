install-tools:
	go install github.com/bufbuild/buf/cmd/buf@v1.59.0

install-deps:
	go mod tidy

clear-protogen:
	rm -rf ./protogen

generate-proto:
	buf dep update \
	&& buf generate --template buf.gen.yaml
