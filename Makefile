proto-fmt:
	find ./proto/ -name "*.proto" | xargs clang-format -i

protoc:
	prototool generate

proto-mock:
	mkdir -p proto/api/v1beta1/mock
	mockgen -source=proto/api/v1beta1/api.pb.go -package=mock -destination=proto/api/v1beta1/mock/api.pb.go
