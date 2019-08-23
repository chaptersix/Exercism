package space

type Planet string

var m = map[Planet]float64{
	"Earth":   31557600,
	"Mercury": 7600543.81992,
	"Venus":   19414149.052176,
	"Mars":    59354032.690079994,
	"Jupiter": 374355659.124,
	"Saturn":  929292362.8848,
	"Uranus":  2651370019.3296,
	"Neptune": 5200418560.032001,
}

//Age returns your age on the given planet
func Age(sec float64, p Planet) float64 {
	return sec / m[p]
}
