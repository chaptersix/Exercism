import java.util.Arrays;
class Triangle {
    private double[] sides;

    Triangle(double side1, double side2, double side3) throws TriangleException {
        sides = new double[] {side1,side2,side3};
        Arrays.sort(sides);
        if  (sides[0] <= 0) {
            throw new TriangleException();
        }
        if ((sides[0] + sides[1] < sides[2])) {
            throw new TriangleException();
        }
    }

    boolean isEquilateral() {
        return sides[0] == sides[2];
    }

    boolean isIsosceles() {
        return (sides[0] == sides[1] || sides[1] == sides[2]);
    }

    boolean isScalene() {
        return !isIsosceles();
    }

}
