.PHONY: image.antlr
image.antlr:
	DOCKER_BUILDKIT=1 docker build \
		-t antlr \
		-f docker/Dockerfile.antlr .


.PHONY: run.antlr
run.antlr.day8:
	docker run -it --rm \
		-v $(PWD)/pkg/day8parser:/mnt/src \
		antlr \
			-no-listener \
			-visitor \
			-Dlanguage=Go \
			-o /mnt/src/ \
			-package day8parser \
			/mnt/src/Expressions.g4

# https://www.honeybadger.io/blog/golang-go-package-management/
.PHONY: go.requirements.compile
go.requirements.compile:
	go mod tidy

.PHONY: go.fmt
go.fmt:
	gofmt -s -w .
