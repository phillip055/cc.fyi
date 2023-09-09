import typer
from decoder import HuffmanTreeDecoder
from encoder import HuffmanTreeEncoder

app = typer.Typer()

@app.command()
def encode(filepath: str, resultpath: str):
    io = open(filepath)
    content = io.read()
    tree = HuffmanTreeEncoder(content)
    result = tree.compress()
    with open(resultpath, "w") as f:
        f.write(dict(tree.freq_table).__str__())
        f.write('\n')
        f.write(result)

@app.command()
def decode(filepath: str, resultpath: str):
    io = open(filepath)
    header = io.readline()
    content = io.readline()
    result = HuffmanTreeDecoder().decompress(eval(header), content)
    with open(resultpath, "w") as f:
        f.write(result)


if __name__ == "__main__":
    app()