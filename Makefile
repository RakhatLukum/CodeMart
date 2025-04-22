PROTO_DIR=proto
GO_OUT=.

proto:
	protoc -I $(PROTO_DIR) \
		--go_opt=module=CodeMart \
		--go-grpc_opt=module=CodeMart \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_DIR)/*.proto
