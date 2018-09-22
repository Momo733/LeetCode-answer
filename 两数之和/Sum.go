package 两数之和

func TwoSum(nums []int, target int) []int {
	var numsMap = map[int]int{}
	for k,v:=range nums {
		numsMap[v]=k
	}
	for i:=0;i<len(nums);i++{
		if v,ok:=numsMap[target-nums[i]];ok{
			if i!=v{
				return []int{i,v}
			}
		}
	}
	return []int{-1,-1}
}
