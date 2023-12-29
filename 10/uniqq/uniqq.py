from collections import Counter
from io import TextIOWrapper
from typing import Annotated
import typer
import sys
import re
from pathlib import Path


def uniqq(filepath: Annotated[str, typer.Argument()] = '-',
          output_path: Annotated[str, typer.Argument()] = None,
          c: bool = typer.Option(default=False),
          d: bool = typer.Option(default=False),
          u: bool = typer.Option(default=False),
          ):
    io = open(filepath) if filepath != "-" else TextIOWrapper(sys.stdin.buffer)
    counter = Counter()
    while line := io.readline():
        counter.update([line])
    output_print = ""
    for k, v in counter.items():
        if d and v == 1:
            continue
        if u and v > 1:
            continue
        if c:
            output_print += str(v) + " "
        output_print += k
    if output_path:
        output_path = Path(output_path)
        output_path.parent.mkdir(parents=True, exist_ok=True)
        output_path.write_text(output_print)
    else:
        print(output_print, end="")


if __name__ == '__main__':
    typer.run(uniqq)
