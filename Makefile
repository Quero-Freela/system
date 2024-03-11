CGO_ENABLED = 1
GOOS = linux
GOARCH = amd64

.DEFAULT_GOAL := all

.PHONY: all
all: build-ext build-client

.PHONY: build-docker
build-docker:
	docker build -t querofreela .

.PHONY: build-ext
build-ext:
	mkdir -p bin && \
	git submodule sync && \
	git submodule update --init --recursive && \
	cd php-go/ext && \
	phpize && \
	./configure --enable-phpgo && \
	make EXTENSION_DIR=$(realpath ../../bin/.) && \
	make EXTENSION_DIR=$(realpath ../../bin/.) install
	cp php-go/ext/modules/phpgo.so ./bin/.

.PHONY: build-client
build-client:
	mkdir -p bin && \
	cd php-public && \
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -buildmode=c-shared -o ../bin/querofreela.so main.go

.PHONY: publish
publish:
	cp -rf ./public/. ./publish
	cp -rf ./php-public/. ./publish
	cp bin/querofreela.so publish/querofreela.so
	cp bin/phpgo.so publish/phpgo.so
	rm -rf publish/main.go
	chmod -R 777 publish
	cd publish && \
	composer install --no-dev --optimize-autoloader
