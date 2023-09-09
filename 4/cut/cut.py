from io import TextIOWrapper
from typing import Annotated
import typer
import sys

def cut(filepath:Annotated[str, typer.Argument()] = None, f: str = typer.Option(), d: str = typer.Option(default="\t")):
    field_indexes = list(f)
    io = open(filepath) if filepath else TextIOWrapper(sys.stdin.buffer)
    delimiter = d
    while line := io.readline():
        cols = line.split(delimiter)
        filtered_line = delimiter.join([cols[int(i)-1] for i in field_indexes])
        print(filtered_line)

if __name__ == '__main__':
  typer.run(cut)
