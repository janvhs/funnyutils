# TODO: Reference https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html

build:
	rm -rf ./dist
	mkdir ./dist
	go build -o ./dist/ ./cmd/...
	# TODO: rename every binary in dist to f<NAME>
