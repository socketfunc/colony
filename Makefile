proto-fmt:
	find ./proto/ -name "*.proto" | xargs clang-format -i

gazelle:
	bazel run //:gazelle
	gazelle update-repos -from_file=go.mod

test:
	bazel test //...
