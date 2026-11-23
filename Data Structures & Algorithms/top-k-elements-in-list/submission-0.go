func topKFrequent(nums []int, k int) []int {
 size := 1000 - (-1000) + 1;
 freqArr := make([]int, size, size);
 for _ , v := range(nums){
  freqArr[v + 1000]++;
 }

 result := make([]int, 0);
 for k > 0 {
 ind := 0;
 for i, v := range(freqArr){
	if v > freqArr[ind]{
		ind  = i;
	}
 }
 result = append(result, ind - 1000);
 freqArr[ind] = 0;
 k--;
 }
 return result;
}
