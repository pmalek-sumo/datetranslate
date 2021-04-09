markdownlint: mdl

mdl:
	mdl --style .markdownlint/style.rb \
			README.md

build:
	go build -v .

test:
	go test -count 1 -v .

lint:
	golangci-lint run \
	--max-same-issues 50 \
	--enable unparam,unconvert,errorlint,lll,wastedassign
