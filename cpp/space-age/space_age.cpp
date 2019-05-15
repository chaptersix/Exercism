#include "space_age.h"

        space_age::space_age::space_age(long sec_in) {
            second = sec_in;
        }

         double space_age::space_age::seconds() const{
            return second;
        }

        double space_age::space_age::on_earth() const {
            return  second/31557600.0 ;
        }

        double space_age::space_age::on_mercury() const {
            return second/7600543.81992;
        }

        double space_age::space_age::on_venus() const {
            return second/19414149.052;
        }

        double space_age::space_age::on_mars() const {
            return second/59354032.69;
        }

        double space_age::space_age::on_jupiter() const {
            return second/374355659.124;
        }

        double space_age::space_age::on_saturn() const {
            return second/929292362.8848;
        }

        double space_age::space_age::on_uranus() const {
            return second/2651370019.3296;
        }

        double space_age::space_age::on_neptune() const {
            return second/5200418560.032;
        }