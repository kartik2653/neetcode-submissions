type NumMatrix struct {
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
m := len(matrix);
n := len(matrix[0]);
prefixSum := make([][]int, m, m);
for i := 0 ; i < m ; i++{
	prefixSum[i] = make([]int, n, n);
}

for i := 0; i < m ; i++{
	rowSum := 0;
	for j := 0 ; j < n ; j++{
		upperRectSum := 0;
		if i > 0 {
			upperRectSum = prefixSum[i-1][j];
		}
		prefixSum[i][j] =  upperRectSum + matrix[i][j] + rowSum;
		rowSum +=  matrix[i][j];
	}
}
return NumMatrix{
	prefixSum : prefixSum,
}; 
}

func (this *NumMatrix) SumRegion(i int, j int, x int, y int) int {
  result := this.prefixSum[x][y];
  if i - 1 >= 0 {
	result = result - this.prefixSum[i-1][y];
  } 
  if j - 1 >= 0{
	result = result - this.prefixSum[x][j-1];
  }

  if i - 1 >=0 && j-1 >=0 {
	result += this.prefixSum[i-1][j-1]
  } 
  return result;
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
