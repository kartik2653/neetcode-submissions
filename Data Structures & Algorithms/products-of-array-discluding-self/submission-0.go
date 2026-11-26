func productExceptSelf(nums []int) []int {
 n := len(nums);
 output := make([]int, n);
 nonZeroProduct := 1;
 zeroCount := 0;
 for _, v := range (nums){
	if v == 0{
		zeroCount++; 
	}else{
     nonZeroProduct = nonZeroProduct * v;
	} 
 }
 if zeroCount > 1 {
	return output;
 }
 if zeroCount == 1{
  for i, v := range(nums){
	if v == 0 {
		output[i] = nonZeroProduct;
	}else{
		output[i] = 0;
	}
  }
 }else{
  for i, v := range(nums){
	
	output[i] = nonZeroProduct / v;
	
  }
 } 
 return output;
}
