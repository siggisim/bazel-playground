.PHONY: generate
generate:
	@bazel run //:gazelle
	@bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: test
test:
	@bazel test //... --test_output=errors

.PHONY: build
build:
	@bazel build //...
