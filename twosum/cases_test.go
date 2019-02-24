package twosum

var testCases = []struct {
	nums	[]int
	target 	int
	result	[]int
}{
	{
		[]int{2, 7, 11, 15}, 
		9,
		[]int{0, 1},
	},
	{
		[]int{3, 4, 5, 2, 7}, 
		6,
		[]int{1, 3},
	},
	{
		[]int{1, 2, 0, 8, 11, 3}, 
		18,
		nil,
	},
	{
		[]int{}, 
		1,
		nil,
	},
}