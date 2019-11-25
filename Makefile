all: test build install clean

.PHONY: build
build:
	@go build kubenetes_deployer

.PHONY: test
test:
	@go test kubenetes_deployer

.PHONY: install
install:
	@go install kubenetes_deployer

.PHONY: clean
clean:
	@rm kubenetes_deployer

.PHONY: release
release:
	$(eval COMMENT := $(shell bash -c 'read -e -p "Comment: " var; echo $$var'))
	@git add --all; \
	 git commit -m "$(COMMENT)"; \
	 git push