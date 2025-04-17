proto:
	rm -rf ./internal/pb/github.com/webbsalad/pvz/pvz_v1/*

	protoc \
		-I . \
		-I vendor.protogen \
		--validate_out="lang=go:./internal/pb" \
		--go_out=./internal/pb \
		--go-grpc_out=./internal/pb \
		./api/pvz/*.proto

	protoc -I . \
		--grpc-gateway_out ./internal/pb \
		--proto_path=vendor.protogen \
        --grpc-gateway_opt generate_unbound_methods=true \
        ./api/pvz/*.proto

proto-deps:
	rm -rf ./vendor.protogen
	mkdir -p vendor.protogen

	git clone https://github.com/googleapis/googleapis.git ./vendor.protogen/googleapis
	mv ./vendor.protogen/googleapis/google/ ./vendor.protogen
	rm -rf ./vendor.protogen/googleapis/

	git clone https://github.com/bufbuild/protoc-gen-validate.git ./vendor.protogen/protoc-gen-validate
	mv ./vendor.protogen/protoc-gen-validate/validate/ ./vendor.protogen
	rm -rf ./vendor.protogen/protoc-gen-validate/

mocks:
	mockgen -source ./internal/repository/item/repository.go -destination ./internal/repository/item/mock/repository.go -package mock_item
	mockgen -source ./internal/repository/user/repository.go -destination ./internal/repository/user/mock/repository.go -package mock_user
	mockgen -source ./internal/repository/pvz/repository.go -destination ./internal/repository/pvz/mock/repository.go -package mock_pvz
