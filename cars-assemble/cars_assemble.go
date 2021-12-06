package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch {
	case speed < 1:
		return 0
	case speed <= 4:
		return 1
	case speed <= 8:
		return 0.9
	default:
		return 0.77
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	perHour := float64(speed) * 221 * SuccessRate(speed)
	return perHour
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	perMinute := CalculateProductionRatePerHour(speed) / 60
	return int(perMinute)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if speed <= 2 {
		return float64(speed * 221)
	} else {
		return limit
	}
}
