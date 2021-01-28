class Luhn:
    def __init__(self, card_num):
        self.card_num = card_num

    def valid(self):
        index = 0
        sum = 0
        for char in reversed(self.card_num):
            if char == ' ':
                continue
            try:
                nr = int(char)
                if index % 2:
                    nr *= 2
                    if nr > 9:
                        nr -= 9
                index += 1
                sum += nr
            except:
                return False

        return index > 1 and not sum % 10
