.PHONY: faqappd

GO_BUILD_ENV := GOOS=linux GOARCH=amd64

VERSION = $(shell cat VERSION)

IMAGE_NAME = faqapp/faqapp
IMAGE_TAG = $(VERSION)

faqappd: clean
	$(GO_BUILD_ENV) go build -v -o $@ ./cmd/$@

clean:
	rm -f faqappd

image: faqappd
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
