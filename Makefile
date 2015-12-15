target=release

build:
	npm run build
	go build -ldflags "-s" -o $(target)/itpkg main.go
	cp -a config templates $(target)/itpkg


clean:
	-rm $(target)


format:
	for f in `find . -type f -iname '*.go'`; do gofmt -w $$f; done





