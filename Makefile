MAKE_PATH=$(GOPATH)/bin:/bin:/usr/bin:/usr/local/bin:$PATH

.PHONY: all
all: clean format gen lint

BUF_VERSION=1.8.0

.PHONY: buf-install
buf-install:
ifeq ($(shell uname -s), Darwin)
	@[ ! -f $(GOPATH)/bin/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Darwin-x86_64 && \
	mv buf $(GOPATH)/bin/buf && \
	chmod +x $(GOPATH)/bin/buf
else
	@[ ! -f $(GOPATH)/bin/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Linux-x86_64 && \
	mv buf $(GOPATH)/bin/buf && \
	chmod +x $(GOPATH)/bin/buf
endif

.PHONY: gen
gen: buf-install
	@$(GOPATH)/bin/buf generate
	@for dir in $(CURDIR)/gen/go/*/; do \
	  cd $$dir && \
	  path=`awk '{ sub(/.*grpc-contracts/, ""); print }' <<< $$dir` && \
	  path=$$path'/' && \
	  path=`echo $$path | sed s/'\/\/'/''/g` && \
	  go mod init github.com/ilkinabd-talk/grpc-contracts$$path && go mod tidy; \
  	done

.PHONY: lint
lint: buf-install
	@$(GOPATH)/bin/buf lint --config buf.yaml

.PHONY: format
format: buf-install
	@$(GOPATH)/bin/buf format


.PHONY: clean
clean:
	@rm -rf ./gen || true
