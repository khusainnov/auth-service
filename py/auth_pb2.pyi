from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class CreateFileResponse(_message.Message):
    __slots__ = ["code", "message"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    code: int
    message: str
    def __init__(self, code: _Optional[int] = ..., message: _Optional[str] = ...) -> None: ...

class DeleteRequest(_message.Message):
    __slots__ = ["username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    def __init__(self, username: _Optional[str] = ...) -> None: ...

class DeleteResponse(_message.Message):
    __slots__ = ["message", "status"]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    message: str
    status: int
    def __init__(self, status: _Optional[int] = ..., message: _Optional[str] = ...) -> None: ...

class FileRequest(_message.Message):
    __slots__ = ["username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    def __init__(self, username: _Optional[str] = ...) -> None: ...

class ResponseAuth(_message.Message):
    __slots__ = ["role_name"]
    ROLE_NAME_FIELD_NUMBER: _ClassVar[int]
    role_name: str
    def __init__(self, role_name: _Optional[str] = ...) -> None: ...

class ResponseFile(_message.Message):
    __slots__ = ["chunks", "code", "username"]
    CHUNKS_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    chunks: _containers.RepeatedScalarFieldContainer[bytes]
    code: int
    username: str
    def __init__(self, code: _Optional[int] = ..., username: _Optional[str] = ..., chunks: _Optional[_Iterable[bytes]] = ...) -> None: ...

class ResponseMsg(_message.Message):
    __slots__ = ["code", "message"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    code: int
    message: str
    def __init__(self, code: _Optional[int] = ..., message: _Optional[str] = ...) -> None: ...

class ResponseName(_message.Message):
    __slots__ = ["name", "patronymic", "surname"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PATRONYMIC_FIELD_NUMBER: _ClassVar[int]
    SURNAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    patronymic: str
    surname: str
    def __init__(self, name: _Optional[str] = ..., surname: _Optional[str] = ..., patronymic: _Optional[str] = ...) -> None: ...

class StatisticsRequest(_message.Message):
    __slots__ = ["username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    def __init__(self, username: _Optional[str] = ...) -> None: ...

class StatisticsResponse(_message.Message):
    __slots__ = ["code", "file_numb", "message", "user_numb"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    FILE_NUMB_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    USER_NUMB_FIELD_NUMBER: _ClassVar[int]
    code: int
    file_numb: int
    message: str
    user_numb: int
    def __init__(self, code: _Optional[int] = ..., message: _Optional[str] = ..., user_numb: _Optional[int] = ..., file_numb: _Optional[int] = ...) -> None: ...

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

class UserFile(_message.Message):
    __slots__ = ["file", "username"]
    FILE_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    file: bytes
    username: str
    def __init__(self, username: _Optional[str] = ..., file: _Optional[bytes] = ...) -> None: ...

class UserRequest(_message.Message):
    __slots__ = ["password", "username"]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    password: str
    username: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...
