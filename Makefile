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
	rm -f ./py/pb/*.pb.go

cleanpy:
	rm -f ./py/*pb2*.py \
	rm -f ./py/*.pyi

m-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' up
m-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' down
d-up:
	docker run --name=gauth -e POSTGRES_PASSWORD='qwerty' -p 5434:5432 -d --rm postgres
d-exec:
	docker exec -it gauth /bin/bash
d-stop:
	docker stop gauth

