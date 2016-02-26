all: goxpath
test: codeclimate

goxpath: *.go
	go build

codeclimate:
	docker run --rm --env CODE_PATH="$(PWD)" \
		--volume "$(PWD)":/code \
		--volume /var/run/docker.sock:/var/run/docker.sock \
		--volume /tmp/cc:/tmp/cc \
		codeclimate/codeclimate validate-config

clean:
	rm -f goxpath

.PHONY: all clean test codeclimate
