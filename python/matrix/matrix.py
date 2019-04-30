class Matrix(object):
    def __init__(self, matrix_string):
        self.matrix_array = []
        for row in matrix_string.split("\n"):
            self.matrix_array.append(
                list(
                    map(int,
                    row.split(" "))))

    def row(self, index):
        return  self.matrix_array[index-1]


    def column(self, index):
        return [column[index-1] for column in self.matrix_array ]

