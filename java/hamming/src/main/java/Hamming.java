import java.util.stream.IntStream;

class Hamming {
    String left;
    String right;

    Hamming(String leftStrand, String rightStrand) {
       if(leftStrand.length() != rightStrand.length()) {
            throw new IllegalArgumentException("leftStrand and rightStrand must be of equal length.");
         }
        left = leftStrand;
        right = rightStrand;
    }

    int getHammingDistance() {
        return (int) IntStream
                .range(0, left.length())
                .filter(x -> left.charAt(x) != right.charAt(x))
                .count();
    }

}