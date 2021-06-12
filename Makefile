.PHONY : all test test-%

all: test

test: test-codewars

test-codewars:
	@ginkgo -r ./codewars
