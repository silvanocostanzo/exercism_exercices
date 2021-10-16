package lasagna

const OvenTime = 40

func RemainingOvenTime(remainingTime int) int {
	return OvenTime - remainingTime
}

func PreparationTime(layers int) int {
	return layers * 2
}

func ElapsedTime(layers int, time int) int {
	return layers*2 + time
}
