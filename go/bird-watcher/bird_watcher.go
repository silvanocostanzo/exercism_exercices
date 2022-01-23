package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	counter := 0
	for _, v := range birdsPerDay {
		counter += v
	}
	return counter
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	counter := 0
	var givenWeek []int
	i := 7 * (week - 1)
	givenWeek = birdsPerDay[i : i+7]
	for i := 0; i < len(givenWeek); i++ {
		counter += givenWeek[i]
	}
	return counter
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = v + 1
		}
	}
	return birdsPerDay
}
