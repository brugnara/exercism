class Matrix:
    def __init__(self, matrix_string):
        self.matrix = []

        for row in matrix_string.split('\n'):
            self.matrix.append(list(map(lambda n: int(n), row.split(' '))))

        self.flippedMatrix = list(map(lambda l: list(l), zip(*self.matrix)))

    def row(self, index):
        return self.matrix[index - 1]

    def column(self, index):
        return self.flippedMatrix[index - 1]
