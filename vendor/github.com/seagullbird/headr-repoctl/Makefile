GOARCH?=amd64
GOOS?=linux
APP?=repoctl
PROJECT?=github.com/seagullbird/headr-repoctl
COMMIT?=$(shell git rev-parse --short HEAD)
PORT?=8687


clean:
	rm -f ${APP}

build: clean
	GOARCH=${GOARCH} GOOS=${GOOS} go build \
	-ldflags "-s -w -X ${PROJECT}/config.PORT=:${PORT}" \
	-o ${APP}

container: build
	docker build -t repoctl:${COMMIT} .

minikube: container
	cat k8s/k8s.yaml | \
		gsed -E "s/\{\{(\s*)\.Commit(\s*)\}\}/$(COMMIT)/g" | \
		gsed -E "s/\{\{(\s*)\.Port(\s*)\}\}/$(PORT)/g" > tmp.yaml
	kubectl apply -f tmp.yaml
	rm -f tmp.yaml
