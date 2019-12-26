PHONY: build-ui build-apiserver


build-ui:
	rm -rf hack/docker/ui && cd ui && npm run build && mv dist/ui ../hack/docker
	cd hack/docker && docker build . -f Dockerfile.ui -t chengtian/koala:v0.0.1

build-apiserver:
	GOOS=linux GOARCH=amd64 go build -o koala ./pkg/cmd/ && mv koala hack/docker/
	cd hack/docker && docker build . -f Dockerfile.apiserver -t chengtian/koala-apiserver:v0.0.1


