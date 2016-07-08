GO_BUILD_ENV := GOOS=linux GOARCH=amd64

faqappd: clean
	$(GO_BUILD_ENV) go build -v -o $@ ./cmd/$@

clean:
	rm -f faqappd

heroku: faqappd
	heroku container:push web
