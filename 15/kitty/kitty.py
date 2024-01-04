import typer
from typing import Annotated
from io import TextIOWrapper
import sys

def kitty(filepath: Annotated[str, typer.Argument()] = '-', n: bool = typer.Option(default=False), filepath0: Annotated[str, typer.Argument()] = None):
    io = open(filepath) if filepath != "-" else TextIOWrapper(sys.stdin.buffer)
    counter = 0
    while line := io.readline():
        counter += 1
        print((str(counter) + "  " if n else "") + line, end="")
    if filepath0:
        io = open(filepath0) if filepath0 != "-" else TextIOWrapper(sys.stdin.buffer)
        while line := io.readline():
            counter += 1
            print((str(counter) + "  " if n else "") + line, end="")

if __name__ == '__main__':
    typer.run(kitty)
