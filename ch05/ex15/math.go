package main

func main() {
}

func max(values ...int) (int, bool) {
	if len(values) == 0 {
		return 0, false
	}
	max := values[0]
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}
	return max, true
}

func min(values ...int) (int, bool) {
	if len(values) == 0 {
		return 0, false
	}
	min := values[0]
	for _, value := range values[1:] {
		if value < min {
			min = value
		}
	}
	return min, true
}

func max2(first int, remaining ...int) int {
	max := first
	for _, value := range remaining {
		if value > max {
			max = value
		}
	}
	return max
}

func min2(first int, remaining ...int) int {
	min := first
	for _, value := range remaining {
		if value < min {
			min = value
		}
	}
	return min
}
