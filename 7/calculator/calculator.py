from typing import Annotated
import typer

precendence = {
    "(": 0,
    ")": 0,
    "-": 1,
    '+': 1,
    "/": 2,
    "*": 2,
}

def calculate(expression:Annotated[str, typer.Argument()]):
    output = []
    operators = []
    tokens = expression.split()
    for token in tokens:
        if token.isnumeric():
            output.append(float(token))
        else:
            if token == "(":
                operators.append(token)
            elif token == ")":
                while len(operators) and operators[-1] != "(":
                    tbp = operators.pop()
                    output.append(tbp)
                operators.pop()
            else:
                if len(operators) and precendence[token] <= precendence[operators[-1]]:
                    output.append(operators.pop())
                operators.append(token)
    while len(operators):
        output.append(operators.pop())

    stack = []
    for o in output:
        if o in precendence:
            if o == "-":
                stack.append(-1 * (float(stack.pop()) - float(stack.pop())))
            elif o == "*":
                stack.append(float(stack.pop()) * float(stack.pop()))
            elif o == "/":
                denom = float(stack.pop())
                num = float(stack.pop())
                stack.append(num / denom)
            else:
                stack.append(float(stack.pop()) + float(stack.pop()))
        else:
            stack.append(o)
    
    result = int(stack[0]) if stack[0] == int(stack[0]) else stack[0]
    print(result)
    return str(result)

if __name__ == '__main__':
    typer.run(calculate)
