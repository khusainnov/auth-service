from __future__ import print_function
import logging
import grpc

from pb import auth_pb2 as pb
from pb import auth_pb2_grpc as pbg

name = "Rustam"
surname = "Khusainov"
username = "khusainnov"
email = "test@gmail.com"
phone = "+71234567890"
password = "qwerty"

u = pb.User(Name=name, Surname=surname, Username=username, Email=email, Phone=phone, Password=password)


def run():
    with grpc.insecure_channel('127.0.0.1:9090') as channel:
        stub = pbg.AuthServiceStub(channel)
        resp = stub.CreateUser(
            pb.User(Name=u.Name, Surname=u.Surname, Username=u.Username, Email=u.Email, Phone=u.Phone, Password=u.Password))
    print(resp.Message)


if __name__ == '__main__':
    logging.basicConfig()
    run()