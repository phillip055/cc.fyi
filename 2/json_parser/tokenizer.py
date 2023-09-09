from token_types import * 
import re

class Tokenizer:
    def __init__(self, s):
        self.tokens = []
        self.s = s
        self.idx = 0

    def opening_array_token_handler(self):
        self.tokens.append(OpeningArrayToken(self.s[self.idx]))
        self.idx += 1
    
    def closing_array_token_handler(self):
        self.tokens.append(ClosingArrayToken(self.s[self.idx]))
        self.idx += 1
    
    def opening_object_token_handler(self):
        self.tokens.append(OpeningObjectToken(self.s[self.idx]))
        self.idx += 1

    def closing_object_token_handler(self):
        self.tokens.append(ClosingObjectToken(self.s[self.idx]))
        self.idx += 1
    
    def colon_token_handler(self):
        self.tokens.append(ColonToken(self.s[self.idx]))
        self.idx += 1

    def comma_token_handler(self):
        self.tokens.append(CommaToken(self.s[self.idx]))
        self.idx += 1
    
    def true_token_handler(self):
        self.tokens.append(BooleanToken(self.s[self.idx: self.idx + 4]))
        self.idx += 4
    
    def false_token_handler(self):
        self.tokens.append(BooleanToken(self.s[self.idx: self.idx + 5]))
        self.idx += 5
    
    def null_token_handler(self):
        self.tokens.append(NullToken(self.s[self.idx: self.idx + 4]))
        self.idx += 4

    def number_token_handler(self):
        remaining_s = self.s[self.idx:]
        res = re.match('[+-]?(?=\.\d|\d)(?:\d+)?(?:\.?\d*)(?:[eE][+-]?\d+)?', remaining_s)
        if res:
            self.idx += res.end() - res.start()
        else:
            raise ValueError('Invalid number token')
            
 
    def string_token_handler(self):
        token = StringToken(self.s[self.idx])
        self.idx += 1
        while self.s[self.idx] != '"':
            if self.s[self.idx] == '\t':
                raise ValueError("Unexpected tab character inside string literal")
            if self.s[self.idx] == "\n":
                raise ValueError("Unexpected line break inside string literal")
            if self.s[self.idx] == "\\":
                if self.s[self.idx+1] in ['"', "\\", "/", "b", "f", "n", "r", "t"]:
                    token.value += self.s[self.idx]
                    token.value += self.s[self.idx+1]
                    self.idx += 2
                    continue
                else:
                    raise ValueError("Invalid escape character")        
            token.value += self.s[self.idx]
            self.idx += 1
        token.value += self.s[self.idx]
        self.tokens.append(token)
        self.idx += 1

    def tokenize(self):
        while self.idx < len(self.s):
            match self.s[self.idx]:
                case "[": self.opening_array_token_handler()
                case "]": self.closing_array_token_handler()
                case "{": self.opening_object_token_handler()
                case "}": self.closing_object_token_handler()
                case ":": self.colon_token_handler()
                case '"': self.string_token_handler()
                case ",": self.comma_token_handler()
                case "t" if len(self.s) > 4 and self.s[self.idx: self.idx + 4] == "true":
                    self.true_token_handler()
                case "f" if len(self.s) > 5 and self.s[self.idx: self.idx + 5] == "false":
                    self.false_token_handler()
                case "n" if len(self.s) > 4 and self.s[self.idx: self.idx + 4] == "null":
                    self.null_token_handler()
                case " ": self.idx += 1
                case "\n": self.idx += 1
                case '-': self.number_token_handler()
                case '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9':
                    self.number_token_handler()
                case _:
                    raise ValueError("Invalid token")
        return self.tokens
