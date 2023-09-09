from huffman_tree import HuffmanTree

class HuffmanTreeDecoder:
    def decompress(self, freq_table, content):
        content = list(content)
        huffmantree = HuffmanTree(freq_table)
        self.tree = huffmantree.tree
        result = ""
        curr = self.tree
        for idx in range(len(content)):
            if content[idx] == '0':
                curr = curr.left
            else:
                curr = curr.right
            if curr.left is None and curr.right is None:
                result += curr.char
                curr = self.tree
        return result

