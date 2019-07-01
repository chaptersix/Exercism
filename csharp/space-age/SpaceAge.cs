using System;

public class SpaceAge
{
    private int second;
    public SpaceAge(int seconds)
    {
        second = seconds;
    }

    public double OnEarth()
    {
        return second / 31557600.0;
    }

    public double OnMercury()
    {
        return second / 7600543.81992;
    }

    public double OnVenus()
    {
        return second / 19414149.052; ;
    }

    public double OnMars()
    {
        return second / 59354032.69;
    }

    public double OnJupiter()
    {
        return second / 374355659.124;
    }

    public double OnSaturn()
    {
        return second / 929292362.8848;
    }

    public double OnUranus()
    {
        return second / 2651370019.3296;
    }

    public double OnNeptune()
    {
        return second / 5200418560.032;
    }
}
