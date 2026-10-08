func sortColors(nums []int) {
 zC := 0;
 oC := 0;
 tC := 0;

 for _, v := range(nums){
	if v == 0 {
		zC++; 
	}
	if v == 1 {
		oC++;
	}
	if v == 2 {
		tC++;
	}
 }

 for i, _ := range(nums){
	if zC != 0 {
		nums[i] = 0;
		zC--;
	}else if(oC != 0){
		nums[i] = 1;
		oC--;
	}else{
		nums[i] = 2;
	}
 }  

 return; 
}
