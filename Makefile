all: goxpath

goxpath: *.go
	go build

clean:
	rm -f goxpath

.PHONY: all clean
