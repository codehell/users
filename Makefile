go-test:
	docker run --rm --interactive --tty \
  --volume $(PWD):/app \
  -w=/app \
  golang:1.14-buster go test -v