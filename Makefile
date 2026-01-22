

all: help

help:
	@echo "Makefile commands:"
	@echo "  make          - run help"
	@echo "  make test     - Run tests for the ipfs-cid binary"
	@echo "  make info     - Show version information"
	@echo "  make clean    - Remove built binaries"
	@echo "  make publish  - Publish the module to Go proxy(after: git tag v0.x.0)"

ipfs-cid: ipfs-cid.go
	go build -o ipfs-cid ipfs-cid.go

test: ipfs-cid info
	./ipfs-cid
	./ipfs-cid ipfs-cid.go
	cat ipfs-cid.go | ./ipfs-cid -stdin

info:
	@echo "version:$$(git describe --tags --match="v[0-9]*" --abbrev=0 HEAD)"
	@echo

clean:
	rm -f ipfs-cid

# https://stackoverflow.com/questions/3867619/how-to-get-last-git-tag-matching-regex-criteria
# git tag v0.x.0
# make publish
publish:
	git push --tags
	GOPROXY=proxy.golang.org go list -m "github.com/lrobot/ipfs-cid@$$(git describe --tags --match="v[0-9]*" --abbrev=0 HEAD)"
publish_by_curl:
	git push --tags
	curl -v "https://pkg.go.dev/github.com/lrobot/ipfs-cid@$$(git describe --tags --match="v[0-9]*" --abbrev=0 HEAD)"

