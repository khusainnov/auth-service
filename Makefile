gengo:
	protoc -I ./proto \
	--go_out=./gen/pb/ --go_opt paths=source_relative \
	--go-grpc_out ./gen/pb/ --go-grpc_opt paths=source_relative \
	./proto/auth.proto

genpy:
	python3 -m grpc_tools.protoc -I./proto \
	--python_out=./py/pb/ --pyi_out=./py/pb --grpc_python_out=./py/pb \
	./proto/auth.proto

cleango:
	rm -f ./py/pb/*.pb.go

cleanpy:
	rm -f ./py/pb/*.py \
	rm -f ./py/pb/*.pyi

