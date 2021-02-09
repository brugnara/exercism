def is_isogram(string):
    seen = {}
    for c in string.lower():
        if c < 'a' or c > 'z':
            continue
        if c in seen:
            return False
        seen[c] = True
    return True
