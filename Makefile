.PHONY : all test test_%

all: test

test: test-codewars
	@echo "** Test finished **"

test-codewars:
	ginkgo -r ./codewars
