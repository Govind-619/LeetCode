import "math"

func angleClock(hour int, minutes int) float64 {
    hour %= 12

    minuteAngle := float64(minutes) * 6.0

    hourAngle := float64(hour)*30.0 + float64(minutes)*0.5

    diff := math.Abs(hourAngle - minuteAngle)

    return math.Min(diff, 360.0-diff)
}