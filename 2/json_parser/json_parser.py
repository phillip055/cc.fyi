from tokenizer import *

class JsonParser:
    def __init__(self, input):
        self.tokens = Tokenizer(input).tokenize()
    
    def parse(self):
        if isinstance(self.tokens[0], OpeningObjectToken):
            self.__parseObject()
        elif isinstance(self.tokens[0], OpeningArrayToken):
            self.__parseArray()
        else:
            raise SyntaxError("Invalid start token")
        if len(self.tokens):
            raise SyntaxError("Invalid end token")
    
    def __parseObject(self):
        if not isinstance(token:=self.tokens.pop(0), OpeningObjectToken):
            raise SyntaxError("Invalid opening object")
        while len(self.tokens) and not isinstance(token := self.tokens.pop(0), ClosingObjectToken):
            if isinstance(token, StringToken):
                if isinstance(token := self.tokens.pop(0), ColonToken):
                    self.__parseValue()
                else:
                    raise SyntaxError("Missing expected colon token")
            else:
                raise SyntaxError("Not a valid key token")
            if isinstance(token := self.tokens[0], CommaToken):
                self.tokens.pop(0)
                if isinstance(token := self.tokens[0], ClosingObjectToken):
                    raise SyntaxError("Extra Comma")
            
    def __parseArray(self):
        if not isinstance(token:=self.tokens.pop(0), OpeningArrayToken):
            raise SyntaxError("Invalid opening array")
        while len(self.tokens) and not isinstance(token := self.tokens.pop(0), ClosingArrayToken):
            self.tokens.insert(0, token)
            self.__parseValue()
            if isinstance(token := self.tokens[0], CommaToken):
                self.tokens.pop(0)
                if isinstance(token := self.tokens[0], ClosingArrayToken):
                    raise SyntaxError("Extra Comma")

    def __parseValue(self):
        val = self.tokens[0]
        if isinstance(val, StringToken) or isinstance(val, BooleanToken) or isinstance(val, NumberToken) or isinstance(val, NullToken):
            self.tokens.pop(0)
        elif isinstance(val, OpeningObjectToken):
            self.__parseObject()
        elif isinstance(val, OpeningArrayToken):
            self.__parseArray()
        else:
            raise SyntaxError("Invalid value token")
