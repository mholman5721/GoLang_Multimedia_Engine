package mathhelper

// ScaleBetween outputs the value of a number scaled between two allowed values
func ScaleBetween(unscaledNum, minAllowed, maxAllowed, min, max float64) float64 {
	return (maxAllowed-minAllowed)*(unscaledNum-min)/(max-min) + minAllowed
}
