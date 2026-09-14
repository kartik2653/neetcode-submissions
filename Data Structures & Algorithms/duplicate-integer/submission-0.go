func hasDuplicate(nums []int) bool {
    freqSet := make(map[int]any);
	for _, val := range(nums){
		_, isPresent := freqSet[val];
		if isPresent {
			return true;
		}
		freqSet[val] = -1;
	} 

	return false;
}
