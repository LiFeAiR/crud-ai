grpc-gen:
# sudo ln -s ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway
	protoc -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--grpc-gateway_out=logtostderr=true:./pkg/server \
	--swagger_out=allow_merge=true,merge_file_name=api:./api/openapi \
	--go_out=plugins=grpc:./pkg/server ./api/grpc/*.proto