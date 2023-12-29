from io import TextIOWrapper
from typing import Annotated
import typer
import sys
import re
from pathlib import Path

def grepp_file(path, pattern, io, i, r):
    while line := io.readline():
        if i:
            match = re.search(pattern, line, flags=re.IGNORECASE)
        else:
            match = re.search(pattern, line)
        if match:
            if r:
                print(str(path) + ":" + line, end="")
            else:
                print(line, end="")

def ggrep(pattern: Annotated[str, typer.Argument()],
          filepath: Annotated[str, typer.Argument()] = None,
          r: bool = typer.Option(default=False),
          i: bool = typer.Option(default=False)
          ):
    if r:
        for path in Path(filepath).rglob('*'):
            if not path.is_file():
                continue
            io = open(path)
            grepp_file(path, pattern, io, i)
    else:
        io = open(filepath) if filepath else TextIOWrapper(sys.stdin.buffer)
        grepp_file(
            path=filepath,
            pattern=pattern,
            io=io,
            i=i,
            r=r
        )

if __name__ == '__main__':
    typer.run(ggrep)
