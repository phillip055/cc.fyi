import collections
from huffman_tree import HuffmanTree

class HuffmanTreeEncoder:
    def __init__(self, s):
        self.s = s
        char_list = list(self.s)
        self.freq_table = collections.Counter(char_list)
        self.code_table = {}

    def __convert(self, char):
        node = self.code_table[char]
        return node.prefix

    def compress(self):
        tree = HuffmanTree(self.freq_table)
        self.code_table = tree.code_table
        r = ""
        for char in list(self.s):
            r += self.__convert(char)
        return r
