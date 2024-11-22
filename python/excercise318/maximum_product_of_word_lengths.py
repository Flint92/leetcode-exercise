from typing import List


def maximum_product_of_word_lengths(words: List[str]) -> int:
    n = len(words)
    bits = [0] * n

    ans = 0
    for i in range(n):
        for ch in words[i]:
            bits[i] |= 1 << (ord(ch) - ord('a'))
        for j in range(0, i):
            if bits[i] & bits[j] == 0:
                prod = len(words[i]) * len(words[j])
                ans = max(ans, prod)
    return ans
