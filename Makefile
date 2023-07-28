# TODO: Reference https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html

build:
	rm -rf ./dist
	mkdir ./dist
	CGO_ENABLED=0 go build -o ./dist/ ./cmd/...
	go run ./tools/prefix.go ./dist/
