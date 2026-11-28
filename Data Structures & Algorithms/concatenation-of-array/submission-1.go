func getConcatenation(nums []int) []int {
    n := len(nums);
    ans := make([]int , 2*n);
    for i := range(2*n){
        ans[i] = nums[i % n];
    }
    return ans;
}
