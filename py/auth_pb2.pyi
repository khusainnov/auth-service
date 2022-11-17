from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ResponseMsg(_message.Message):
    __slots__ = ["Code", "Message"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    Code: int
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    Message: str
    def __init__(self, Code: _Optional[int] = ..., Message: _Optional[str] = ...) -> None: ...

class User(_message.Message):
    __slots__ = ["Email", "Name", "Password", "Patronymic", "Surname", "Username"]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    Email: str
    NAME_FIELD_NUMBER: _ClassVar[int]
    Name: str
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PATRONYMIC_FIELD_NUMBER: _ClassVar[int]
    Password: str
    Patronymic: str
    SURNAME_FIELD_NUMBER: _ClassVar[int]
    Surname: str
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    Username: str
    def __init__(self, Username: _Optional[str] = ..., Name: _Optional[str] = ..., Surname: _Optional[str] = ..., Patronymic: _Optional[str] = ..., Email: _Optional[str] = ..., Password: _Optional[str] = ...) -> None: ...

class UserRequest(_message.Message):
    __slots__ = ["Password", "Username"]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    Password: str
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    Username: str
    def __init__(self, Username: _Optional[str] = ..., Password: _Optional[str] = ...) -> None: ...
