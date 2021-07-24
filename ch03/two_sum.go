package ch03

func TwoSum(nums []int, target int) []int  {
	d := make(map[int]int, len(nums))
	for i, n := range nums{
		if m, ok := d[target - n]; ok {
			return []int{m, i}
		}
		d[n] = i
	}

	return []int{}
}
