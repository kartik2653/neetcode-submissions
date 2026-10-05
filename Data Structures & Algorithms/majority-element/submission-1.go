func majorityElement(nums []int) int {
  mE := nums[0];
  vote := 0;
  for i := 0 ; i < len(nums) ; i++{
	if nums[i] == mE{
		vote++;
	}else{
		vote--;
	}
	if vote <= 0 {
		mE = nums[i];
		vote = 0;
	}
  }

  return mE;

}
