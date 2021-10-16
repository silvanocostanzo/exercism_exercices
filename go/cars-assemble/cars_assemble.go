package cars

const carsPerHour = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return successRate(speed) * carsPerHour * float64(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed)) / 60
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	switch speed {
	case 1, 2, 3, 4:
		return 1
	case 5, 6, 7, 8:
		return 0.9
	case 9, 10:
		return 0.77
	default:
		return 0
	}
}
