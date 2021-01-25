blank:
	@echo "Specify command explicitly"

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

# =================================================

VENDOR:=dev.dotsenko
APP_NAME:=aoc2020

DAY?=11
JDK_VERSION:=11
JDK_VENDOR:=openjdk
MVN_VERSION:=3.6.3
MVN_IMAGE:=maven:$(MVN_VERSION)-$(JDK_VENDOR)-$(JDK_VERSION)
ARGS?=


mvn.repo:
	docker volume create --name maven-repo


_PWD:=/usr/src/this
MVN:=docker run -it --rm \
              		-v "$(PWD)/day$(DAY)":$(_PWD) \
              		-v maven-repo:/root/.m2 \
              		-w $(_PWD) \
              		$(MVN_IMAGE)

mvn.create: mvn.repo
	$(MVN) archetype:generate \
		-DgroupId=$(VENDOR) \
		-DartifactId=$(APP_NAME) \
		-DarchetypeArtifactId=maven-archetype-quickstart


mvn: mvn.repo
	$(MVN) mvn $(ARGS)

.PHONY: mvn.sh
mvn.sh: mvn.repo
	$(MVN) bash
