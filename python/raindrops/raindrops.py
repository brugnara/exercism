def convert(number):
    ret = ''
    ret += 'Pling' if not number % 3 else ''
    ret += 'Plang' if not number % 5 else ''
    ret += 'Plong' if not number % 7 else ''

    return ret if ret != '' else str(number)
