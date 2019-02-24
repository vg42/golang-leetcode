package twosum

func twoSum(nums []int, target int) []int {
	mapDiff := make(map[int]int)

	for i, value := range nums {
		if j, ok := mapDiff[value]; ok {
			return []int{j, i}
		}
		diff := target - value
		mapDiff[diff] = i
	}

    return  nil
}