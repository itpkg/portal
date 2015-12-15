target=release

build:
	npm run build
	go build -ldflags "-s" -o $(target)/itpkg main.go


clean:
	-rm $(target)





