PRODUCT := urlshortener
REPOROOT := gustavoteixeira8
PLATFORM := linux/amd64
PWD := $(shell pwd)
DOCKER_CONTAINERS_IDS := $(shell docker ps -aq)
STACK := urlshortener

product-api:
	$(eval PRODUCT := api)

compose-stop: ; docker stop ${DOCKER_CONTAINERS_IDS}
compose-down: ; docker compose down
compose-up: ; docker compose up -d
stop-api: product-api ; docker stop ${STACK}-${PRODUCT}-1
init-env: ; make compose-stop ; make img-all ; make compose-down ; make compose-up ; clear
init-env-dev: ; make compose-stop ; make compose-up ; make stop-api ; clear
run-api: ; clear ; go run ./cmd/api/main.go

docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deploy/$(PRODUCT)/main \
	./cmd/${PRODUCT}/main.go && \
	cd deploy/$(PRODUCT) && \
	docker build --platform $(PLATFORM) -t $(REPOROOT)/$(STACK)-$(PRODUCT):latest .

docker-push: 
	docker push  $(REPOROOT)/$(STACK)-$(PRODUCT):latest

img-api: product-api docker-build

img-all: 
	make img-api
