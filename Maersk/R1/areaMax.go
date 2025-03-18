package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

}

func diff(i, j int) int {

	if i-j > 0 {
		return i - j
	}
	return j - i

}

func min(i, j int) int {

	if i > j {
		return j
	}
	return i

}

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		maxArea = max(maxArea, min(height[left], height[right])*diff(left, right))
		// move inside on the side where the height is lower
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
