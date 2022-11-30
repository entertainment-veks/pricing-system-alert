package service

func percentageChange(old, new float64) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}
