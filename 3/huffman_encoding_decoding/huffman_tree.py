from __future__ import annotations
from dataclasses import dataclass
import heapq

@dataclass
class HuffmanTreeNode:
    char: str
    freq: int
    right: HuffmanTreeNode = None
    left: HuffmanTreeNode = None
    def __lt__(self, other):
        return self.freq < other.freq

@dataclass
class CodeTableInfo:
    bits: int
    prefix: str

class HuffmanTree:
    def __init__(self, freq_table):
        arr = []
        for key in freq_table:
            arr.append(HuffmanTreeNode(key, freq_table[key]))
        heapq.heapify(arr)
        queue = arr
        while len(queue) > 2:
            first, second = queue.pop(0), queue.pop(0)
            heapq.heappush(queue, HuffmanTreeNode(None, first.freq + second.freq, first, second))
        first, second = queue.pop(0), queue.pop(0)
        dummyHead = HuffmanTreeNode(None, first.freq + second.freq, first, second)
        code_table = {}
        def dfs(node, prefix=""):
            if node == None:
                return
            if node.char:
                code_table[node.char] = CodeTableInfo(len(prefix), prefix)
            dfs(node.left, prefix+"0")
            dfs(node.right, prefix+"1")
        dfs(dummyHead)
        self.tree = dummyHead
        self.code_table = code_table

