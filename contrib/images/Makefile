DOCKER = $(shell which docker)

.PHONY: xchains-api xchains-api-rmi

xchains-api:
	$(DOCKER) build --tag scalarorg/xchains-api -f xchains-api/Dockerfile \
				$(shell git rev-parse --show-toplevel)

xchains-api-rmi:
	$(DOCKER) rmi scalarorg/xchains-api 2>/dev/null; true
