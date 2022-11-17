from __future__ import print_function
import logging
import grpc

import auth_pb2 as pb
import auth_pb2_grpc as pbg

name = "ANy"
surname = "Name"
username = "khusainnov"
email = "test@gmail.com"
patronymic = "fan"
password = "qwerty"


def create_user(username, name, surname, patronymic, email, password):
    with grpc.insecure_channel('localhost:9090') as channel:
        stub = pbg.AuthServiceStub(channel)
        resp = stub.CreateUser(
            pb.User(username=username, name=name, surname=surname, patronymic=patronymic, email=email,
                    password=password))
    print(resp.Message)


def handlers():
    create_user(username, name, surname, patronymic, email, password)


if __name__ == '__main__':
    # logging.basicConfig()
    handlers()
    create_user(username, name, surname, patronymic, email, password)
