package mid

import "math"

func bin(arr []int, val int) int {
	low, height, mid := 0, len(arr)-1, 0
	for low <= height {
		mid = int(math.Round(float64(low) + float64(height)))
		guess := arr[mid]
		if guess == val {
			return mid
		}
		if guess > val {
			height = mid - 1
		}
		if guess < val {
			low = mid + 1
		}
	}
	return -1
}
