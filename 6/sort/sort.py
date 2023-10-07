from io import TextIOWrapper
from typing import Annotated
import typer
import sys

def sort(filepath:Annotated[str, typer.Argument()] = None, u: bool = typer.Option(default=False), algorithm: bool = typer.Option(default=False)):
    io = open(filepath) if filepath else TextIOWrapper(sys.stdin.buffer)
    words = io.readlines()
    words = list(map(lambda x: x.strip(), words))
    if u: words = list(set(words))
    words.sort()
    for word in words:
        print(word)

if __name__ == '__main__':
  typer.run(sort)
