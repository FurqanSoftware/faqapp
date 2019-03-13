.PHONY: faqappd

GO_BUILD_ENV := GOOS=linux GOARCH=amd64

VERSION = $(shell git describe --tags)

IMAGE_NAME = faqapp/faqapp
IMAGE_TAG = $(VERSION:v%=%)

faqappd: clean
	$(GO_BUILD_ENV) go build -mod=vendor -v -o $@ ./cmd/$@

clean:
	rm -f faqappd

image: faqappd
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
