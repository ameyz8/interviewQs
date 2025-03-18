/*

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

*/

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
