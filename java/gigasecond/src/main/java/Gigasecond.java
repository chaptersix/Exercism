import java.time.LocalDate;
import java.time.LocalDateTime;

class Gigasecond {
    private LocalDateTime localdate;
    private static final int Gigasecond = 1000000000;

    Gigasecond(LocalDate birthDate) {
        localdate = birthDate.atStartOfDay();
    }

    Gigasecond(LocalDateTime birthDateTime) {
        localdate = birthDateTime;
    }

    LocalDateTime getDateTime() {
        return localdate.plusSeconds(Gigasecond);
    }

}
