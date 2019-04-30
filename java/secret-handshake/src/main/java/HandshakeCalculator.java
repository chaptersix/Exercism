import java.util.List;
import java.util.ArrayList;
import java.util.Collections;

class HandshakeCalculator {
    private static int WINK = 0b1;
    private static  int DOUBLEBLINK = 0b10;
    private static int CLOSEYOUREYES = 0b100;
    private static int JUMP = 0b1000;
    private static int REVERSE = 0b10000;

    List<Signal> calculateHandshake(int number) {
       ArrayList<Signal> handshake = new ArrayList<Signal>();

        if ((number & WINK) > 0) {
            handshake.add(Signal.WINK);
        }
        if ((number & DOUBLEBLINK) > 0) {
            handshake.add(Signal.DOUBLE_BLINK);
        }
        if ((number & CLOSEYOUREYES) > 0) {
            handshake.add(Signal.CLOSE_YOUR_EYES);
        }
        if ((number & JUMP) > 0) {
            handshake.add(Signal.JUMP);
        }
        if ((number & REVERSE) > 0) {
            Collections.reverse(handshake);
        }
    return handshake;
    }

}
