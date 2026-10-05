import ("slices")
func removeElement(nums []int, val int) int {
  slices.Sort(nums);
  count := 0;
  for _, v  := range(nums){
	if v != val {
		count++;
	}
  }
  pivot := count;
 for i := 0 ; i < count ; i++{
	if nums[i] == val{
		nums[i] = nums[pivot];
		nums[pivot] = val;
		pivot++;
		i--;
	}
 }
 return count;

}
