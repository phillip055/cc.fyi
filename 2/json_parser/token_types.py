from dataclasses import dataclass

@dataclass
class Token:
    value: str

@dataclass
class OpeningObjectToken(Token):
    pass

@dataclass
class ClosingObjectToken(Token):
    pass

@dataclass
class StringToken(Token):
    pass

@dataclass
class WhitespaceToken(Token):
    pass

@dataclass
class CommaToken(Token):
    pass

@dataclass
class ColonToken(Token):
    pass

@dataclass
class NumberToken(Token):
    pass

@dataclass
class BooleanToken(Token):
    pass

@dataclass
class NullToken(Token):
    pass

@dataclass
class OpeningArrayToken(Token):
    pass

@dataclass
class ClosingArrayToken(Token):
    pass
