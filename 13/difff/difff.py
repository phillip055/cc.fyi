from typing import Annotated
import typer

def lowest_common_subsequence(lines1: [str], lines2: [str]):
    table = [[0 for _ in range(len(lines2) + 1)] for _ in range(len(lines1) + 1)]

    for i in range(len(lines1)):
        for j in range(len(lines2)):
            if lines1[i] == lines2[j]:
                table[i + 1][j + 1] = table[i][j] + 1
            else:
                table[i + 1][j + 1] = max(table[i][j + 1], table[i + 1][j])
    lcs = []
    i = len(lines1)
    j = len(lines2)
    while i > 0 and j > 0:
        if lines1[i - 1] == lines2[j - 1]:
            lcs.append(lines1[i - 1])
            i -= 1
            j -= 1
        elif table[i][j - 1] > table[i - 1][j]:
            j -= 1
        else:
            i -= 1

    return lcs[::-1]

def difference(
        common_subsequence: [str],
        lines1: [str],
        lines2: [str]
    ):
    lines1_diff = []
    lines2_diff = []
    for i in range(len(lines1)):
        if lines1[i] not in common_subsequence:
            lines1_diff.append(lines1[i])
    for i in range(len(lines2)):
        if lines2[i] not in common_subsequence:
            lines2_diff.append(lines2[i])
    return lines1_diff, lines2_diff

def format_difference(lines1_diff: [str], lines2_diff: [str]):
    result = []
    idx1, idx2 = 0, 0
    while idx1 < len(lines1_diff) or idx2 < len(lines2_diff):
        if idx1 < len(lines1_diff):
            result.append(
                "< " + lines1_diff[idx1]
            )
        if idx2 < len(lines2_diff):
            result.append(
                "> " + lines2_diff[idx2]
            )
        idx1 += 1
        idx2 += 1
    return result

def difff(filepath1:Annotated[str, typer.Argument()], filepath2:Annotated[str, typer.Argument()]):
    with open(filepath1, "r") as f1, open(filepath2, "r") as f2:
        lines1 = f1.readlines()
        lines2 = f2.readlines()
    common_subsequence = lowest_common_subsequence(lines1, lines2)
    lines1_diff, lines2_diff = difference(common_subsequence, lines1, lines2)
    result = format_difference(lines1_diff, lines2_diff)
    for res in result:
        print(res, end="")

if __name__ == '__main__':
    typer.run(difff)
