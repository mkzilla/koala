PHONY: build-ui build-apiserver


build-ui:
	rm -rf uip && cd ui && npm run build && mv dist/ui ../uip
	docker build . -f Dockerfile.ui -t chengtian/koala:v0.0.1

build-apiserver:
	GOOS=linux GOARCH=amd64 go build -o koala ./pkg/cmd/
	docker build . -f Dockerfile.apiserver -t chengtian/koala-apiserver:v0.0.1


