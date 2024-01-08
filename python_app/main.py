import heapq
import time

class HuffmanNode:
    def __init__(self, char, freq):
        self.char = char
        self.freq = freq
        self.left = None
        self.right = None

    def __lt__(self, other):
        return self.freq < other.freq

def build_huffman_tree(text):
    frequency = {}
    for char in text:
        if char in frequency:
            frequency[char] += 1
        else:
            frequency[char] = 1

    heap = [HuffmanNode(char, freq) for char, freq in frequency.items()]
    heapq.heapify(heap)

    while len(heap) > 1:
        left = heapq.heappop(heap)
        right = heapq.heappop(heap)
        merge_node = HuffmanNode(None, left.freq + right.freq)
        merge_node.left = left
        merge_node.right = right
        heapq.heappush(heap, merge_node)

    return heap[0]

def build_huffman_codes(root, current_code="", code_dict=None):
    if code_dict is None:
        code_dict = {}

    if root is not None:
        if root.char is not None:
            code_dict[root.char] = current_code
        build_huffman_codes(root.left, current_code + "0", code_dict)
        build_huffman_codes(root.right, current_code + "1", code_dict)

    return code_dict

def huffman_compress(text):
    root = build_huffman_tree(text)
    huffman_codes = build_huffman_codes(root)
    
    compressed_text = ''.join([huffman_codes[char] for char in text])
    
    return compressed_text, huffman_codes

def print_huffman_codes(huffman_codes):
    print("Huffman Codes:")
    for char, code in sorted(huffman_codes.items()):
        print(f"{char}: {code}")

# Example usage
original_text = "BCCABBDDAECCBBAEDDCC"

start_time = time.time()

compressed_text, huffman_codes = huffman_compress(original_text)

end_time = time.time()

execution_time_microseconds = (end_time - start_time) * 1e6

print(f"Execution Time: {execution_time_microseconds:.2f} microseconds")
print_huffman_codes(huffman_codes)
