
class NaturalNumber {
    private  int num;
    public NaturalNumber(int numIn) {
        if (numIn <= 0) {
            throw new  IllegalArgumentException("You must supply a natural number (positive integer)");
        }
        this.num = numIn;
    }

    public Classification getClassification() {
        int factorSum = 0;
        for ( int i = 1;  i < this.num; i++) {
            if (this.num % i == 0) {
                factorSum += i;
            }
        }
        if (factorSum == this.num) {
            return Classification.PERFECT;
        }
       if (factorSum > this.num) {
            return Classification.ABUNDANT;
        }
            return Classification.DEFICIENT;
    }

}
