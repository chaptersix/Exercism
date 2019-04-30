class Matrix {
  private int[][] matrix;

  Matrix(String matrixAsString) {
    String[] rowTemp = matrixAsString.split("\n");
    this.matrix = new int[rowTemp.length][];
    for (int i = 0; i < rowTemp.length; i++) {
      String[] stringElements = rowTemp[i].split(" ");
      int[] elements = new int[stringElements.length];
      for (int j = 0; j < stringElements.length; j++) {
        elements[j] = Integer.parseInt(stringElements[j]);
      }
      this.matrix[i] = elements;
    }
  }

  int[] getRow(int rowNumber) {
    return this.matrix[rowNumber];
  }

  int[] getColumn(int columnNumber) {
    int[] col = new int[this.matrix.length];
    for (int i = 0; i < this.matrix.length; i++) {
      col[i] = this.matrix[i][columnNumber];
    }
    return col;
  }
}
