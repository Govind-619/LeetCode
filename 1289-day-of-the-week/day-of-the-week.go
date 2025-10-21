func dayOfTheWeek(day int, month int, year int) string {
    // Sakamoto's month offsets (for months 1..12)
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	y := year
	if month < 3 {
		y -= 1
	}
	// 0 = Sunday, 1 = Monday, ..., 6 = Saturday
	w := (y + y/4 - y/100 + y/400 + t[month-1] + day) % 7
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return weekdays[w]
    
}