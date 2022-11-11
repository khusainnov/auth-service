gengo:
	protoc -I ./proto \
	--go_out=./gen/pb/ --go_opt paths=source_relative \
	--go-grpc_out ./gen/pb/ --go-grpc_opt paths=source_relative \
	./proto/auth.proto

genpy:
	python3 -m grpc_tools.protoc -I./proto \
	--python_out=./py/ --pyi_out=./py/ --grpc_python_out=./py/ \
	./proto/auth.proto

cleango:
	rm -f ./pb/*.pb.go

cleanpy:
	rm -f ./py/*.py \
	rm -f ./py/*.pyi

