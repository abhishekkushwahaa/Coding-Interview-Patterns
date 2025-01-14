package Two_Pointers

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	var left int
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	fmt.Println(nums)
}
