package prefix_sum

type NumArray struct {
	nums  []int
	cache []int
}

func Constructor(nums []int) NumArray {
	sum := 0
	cache := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		cache[i] = sum
	}

	return NumArray{
		nums:  nums,
		cache: cache,
	}
}

func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.cache[right]
	}

	return n.cache[right] - n.cache[left-1]
}
