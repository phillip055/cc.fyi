from io import TextIOWrapper
from typing import Annotated
import typer
import sys

def wwcc(filepath:Annotated[str, typer.Argument()] = None, c:bool=False, l:bool=False, m:bool=False, w:bool=False):
    io = open(filepath) if filepath else TextIOWrapper(sys.stdin.buffer)
    line_count = 0
    word_count = 0
    letter_count = 0
    while line := io.readline():
        line_count += 1
        words = line.split()
        word_count += len(words)
        letter_count += len(line.encode('utf-8'))
    
    source = filepath or ""
    if not c and not l and not m and not w:
       print(" "*3, line_count, word_count, letter_count, source)
    if c:
       print(" ", 0, source)
    if l:
       print(" ", line_count, source)
    if m:
       print(" ", letter_count, source)
    if w:
       print(" ", word_count, source)

if __name__ == '__main__':
  typer.run(wwcc)
