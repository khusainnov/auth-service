from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ResponseDoc(_message.Message):
    __slots__ = ["document"]
    DOCUMENT_FIELD_NUMBER: _ClassVar[int]
    document: bytes
    def __init__(self, document: _Optional[bytes] = ...) -> None: ...

class ResponseMsg(_message.Message):
    __slots__ = ["code", "message"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    code: int
    message: str
    def __init__(self, code: _Optional[int] = ..., message: _Optional[str] = ...) -> None: ...

class ResponseToken(_message.Message):
    __slots__ = ["token"]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    token: str
    def __init__(self, token: _Optional[str] = ...) -> None: ...

class ResponseWork(_message.Message):
    __slots__ = ["code", "file"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    FILE_FIELD_NUMBER: _ClassVar[int]
    code: int
    file: bytes
    def __init__(self, code: _Optional[int] = ..., file: _Optional[bytes] = ...) -> None: ...

class User(_message.Message):
    __slots__ = ["email", "name", "password", "patronymic", "surname", "username"]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PATRONYMIC_FIELD_NUMBER: _ClassVar[int]
    SURNAME_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    email: str
    name: str
    password: str
    patronymic: str
    surname: str
    username: str
    def __init__(self, username: _Optional[str] = ..., name: _Optional[str] = ..., surname: _Optional[str] = ..., patronymic: _Optional[str] = ..., email: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class UserRequest(_message.Message):
    __slots__ = ["password", "username"]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    password: str
    username: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class UserWork(_message.Message):
    __slots__ = ["file", "username"]
    FILE_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    file: bytes
    username: str
    def __init__(self, username: _Optional[str] = ..., file: _Optional[bytes] = ...) -> None: ...
