import java.lang.Math;

class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
       int re = 0;
       for (int i = 1; i <=  input; i++) {
           re += i;
       }
       return  (int) Math.pow(re,2);
    }

    int computeSumOfSquaresTo(int input) {
        int re = 0;
        for (int i = 1; i <=  input; i++) {
            re += (int) Math.pow(i,2);
        }
        return re;
    }

    int computeDifferenceOfSquares(int input) {
        return computeSquareOfSumTo(input) - computeSumOfSquaresTo(input);
    }

}
