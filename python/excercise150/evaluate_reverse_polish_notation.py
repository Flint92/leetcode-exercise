
def evaluate_reverse_polish_notation(tokens):
    stack = []
    for token in tokens:
        if token in '+-*/':
            r = stack.pop()
            l = stack.pop()
            stack.append(str(int(eval(l + token + r))))
        else:
            stack.append(token)
    return int(stack[0])
