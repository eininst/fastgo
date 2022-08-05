build-base:
	docker build -f build/DockerfileBase -t fastgo-base .

build-api:
	docker build -t $(APP)-api --build-arg APP="run_api.go" .

.PHONY: build-base