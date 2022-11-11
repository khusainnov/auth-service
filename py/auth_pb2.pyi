from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class SignUpResponse(_message.Message):
    __slots__ = ["Code", "Message"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    Code: int
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    Message: str
    def __init__(self, Code: _Optional[int] = ..., Message: _Optional[str] = ...) -> None: ...

class User(_message.Message):
    __slots__ = ["Email", "Gender", "Name", "Password", "Phone", "Surname", "Username"]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    Email: str
    GENDER_FIELD_NUMBER: _ClassVar[int]
    Gender: str
    NAME_FIELD_NUMBER: _ClassVar[int]
    Name: str
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    Password: str
    Phone: str
    SURNAME_FIELD_NUMBER: _ClassVar[int]
    Surname: str
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    Username: str
    def __init__(self, Name: _Optional[str] = ..., Surname: _Optional[str] = ..., Username: _Optional[str] = ..., Email: _Optional[str] = ..., Phone: _Optional[str] = ..., Gender: _Optional[str] = ..., Password: _Optional[str] = ...) -> None: ...
