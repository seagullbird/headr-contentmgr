GOARCH?=amd64
GOOS?=linux
APP?=contentmgr
PROJECT?=github.com/seagullbird/headr-contentmgr
COMMIT?=$(shell git rev-parse --short HEAD)
PORT?=8688


clean:
	rm -f ${APP}

build: clean
	GOARCH=${GOARCH} GOOS=${GOOS} go build \
	-ldflags "-s -w -X ${PROJECT}/config.PORT=:${PORT}" \
	-o ${APP}

container: build
	docker build -t sitemgr:${COMMIT} .

minikube: container
	cat k8s/k8s.yaml | \
		gsed -E "s/\{\{(\s*)\.Commit(\s*)\}\}/$(COMMIT)/g" | \
		gsed -E "s/\{\{(\s*)\.Port(\s*)\}\}/$(PORT)/g" > tmp.yaml
	kubectl apply -f tmp.yaml
	rm -f tmp.yaml
