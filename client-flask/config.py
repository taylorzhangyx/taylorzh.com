import os


class Config:
    SECRET_KEY = (
        os.environ.get("SECRET_KEY")
        or b"\x8b\xd7\xabf\x8c\xf9/e\x18t>\xd4\xde?\x1b\x0bN\xda\xeco6\xa7u]"
    )
