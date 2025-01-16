def add_strings(s1: str, s2: str) -> str:
    i, j = len(s1)-1, len(s2)-1
    carry = 0
    result = []

    while i >= 0 or j >= 0 or carry != 0:
        digit1 = int(s1[i]) if i >= 0 else 0
        digit2 = int(s2[j]) if j >= 0 else 0

        total = digit1 + digit2 + carry
        result.append(str(total % 10))
        carry = total // 10

        i -= 1
        j -= 1

    return ''.join(result[::-1])
