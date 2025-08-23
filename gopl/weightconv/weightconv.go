package weightconv

const (
	ZeroKg float64 = 0
)

func KgToLb(k float64) float64 { return k * 2.205 }
func LbToKg(l float64) float64 { return l / 2.205 }
