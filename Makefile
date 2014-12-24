all: bin bin/ogipa

bin/ogipa: $(wildcard src/*.go)
	cd src && go build -p $(shell nproc) -o ../$@

bin:
	mkdir -p bin

clean:
	rm -f bin/ogipo

re: clean all
