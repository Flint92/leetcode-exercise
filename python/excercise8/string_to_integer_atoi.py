from numpy.fft.helper import ifftshift


def string_to_integer_atoi(s: str) -> int:
    s = s.strip()
    if not s:
        return 0

    result = 0
    index = 0
    sign = 1

    int_max, int_min = 2 ** 31 -1, - 2 ** 31

    if s[index] == '-' or s[index] == '+':
        sign = 1 if s[index] == '+' else -1
        index = 1


    for c in s[index:]:
        digit = ord(c) - ord('0')
        if digit < 0 or digit > 9:
            break

        if (result > int_max // 10) or (result == int_max // 10 and digit > int_max % 10):
            return int_max if sign == 1 else int_min
        result = result * 10 + digit

    return result * sign
