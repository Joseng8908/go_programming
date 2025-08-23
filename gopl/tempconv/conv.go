package tempconv

func CToF(c float64) float64 { return (c*9/5 + 32) }

func FToC(f float64) float64 { return (f - 32) * 5 / 9 }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
