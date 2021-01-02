image.antlr:
	DOCKER_BUILDKIT=1 docker build \
		-t antlr \
		-f docker/Dockerfile.antlr .
