package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	//panic("Please implement the TotalBirdCount() function")
	c := 0
	for _, v := range birdsPerDay {
		c += v
	}
	return c
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	//panic("Please implement the BirdsInWeek() function")
	return TotalBirdCount(birdsPerDay[(week-1)*7 : week*7])

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	//panic("Please implement the FixBirdCountLog() function")
	for c := 0; c < len(birdsPerDay); c += 2 {
		birdsPerDay[c] += 1
	}
	return birdsPerDay
}
