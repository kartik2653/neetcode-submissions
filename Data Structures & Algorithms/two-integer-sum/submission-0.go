func twoSum(nums []int, target int) []int {
  store := make(map[int]int);
  result := []int{-1,-1};
  for ind,val := range(nums){
	remainder := target  - val; 
	prevInd, isPresent := store[remainder];
	if isPresent {
		result[0] = prevInd;
		result[1] = ind;  
	}else{
		store[val] = ind
	}
  }
  return result;  
}
