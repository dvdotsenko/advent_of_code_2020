.PHONY: image.antlr
image.antlr:
	DOCKER_BUILDKIT=1 docker build \
		-t antlr \
		-f docker/Dockerfile.antlr .


.PHONY: run.antlr
run.antlr.day7:
	docker run -it --rm \
		-v $(PWD)/pkg/day7parser:/mnt/src \
		antlr \
			-no-listener \
			-visitor \
			-Dlanguage=Go \
			-o /mnt/src/ \
			-package day7parser \
			/mnt/src/Expressions.g4

