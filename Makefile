.PHONY: Golang Library

all: tidy update

tidy:
	@go mod tidy

update:
	@go get -u all

help:
	@echo "make tidy - 整理依赖"
	@echo "make update - 更新依赖"
