func sortArray(nums []int) []int {
  mergeSort(nums, 0, len(nums) - 1);
  return nums;  
}

func mergeSort(nums[]int , start int, end int ) {
	n := end - start + 1;
	if n == 1 {
		return;
	}

	mid := (end + start)/2;
	mergeSort(nums, start, mid);
	mergeSort(nums, mid + 1 , end);

	sortedArr := make([]int, n);
	i := start;
	j := mid + 1;
	k := 0;

	for i <= mid && j <= end{
		if nums[i] <= nums[j]{
			sortedArr[k] = nums[i];
			i++;
		}else{
			sortedArr[k] = nums[j];
			j++;
		}  
		k++;
	}

	for i <= mid {
	 sortedArr[k] = nums[i];	
	 k++;
	 i++;
	} 

	for j <= end{
	 sortedArr[k] = nums[j];
	 k++;
	 j++;	
	}

	for ind, v := range(sortedArr){
     nums[ind + start] = v;
	}

	return;

}
