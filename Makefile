proto-fmt:
	find ./proto/ -name "*.proto" | xargs clang-format -i

gazelle:
	bazel run //:gazelle

update-repos:
	bazel run //:gazelle -- update-repos -from_file=go.mod

test:
	bazel test //...

protoc:
	prototool generate
