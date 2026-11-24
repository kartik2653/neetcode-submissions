type Solution struct{}

func (s *Solution) Encode(strs []string) string {
 encoded := "";
 totalStrings := len(strs);
 encoded = encoded + IntToString(totalStrings);
 for _, v := range(strs){
	encoded = encoded + "-" + IntToString(len(v));
 }
 encoded = encoded + "-";
 for _, v := range(strs){
	encoded = encoded + v;
 }
 fmt.Println(encoded);
 return encoded;
}

func (s *Solution) Decode(encoded string) []string {
 //"2-15-9-HelloWorld----"
 //["2","5","9", "Hello", "World", "","","",""]
 splitAtSeparator := strings.Split(encoded, "-");
 totalStrings := StringToInt(splitAtSeparator[0]);
 result := make([]string,totalStrings);
 wordsLength := make([]int, totalStrings);
 headerLength := len(splitAtSeparator[0]) + totalStrings + 1;
 for i := 1 ; i <= totalStrings ; i++{
	headerLength += len(splitAtSeparator[i]);
	wordsLength[i - 1] =  StringToInt(splitAtSeparator[i]);
 }

 content := encoded[headerLength:]
 pivot := 0;
 for i, v := range(wordsLength){
  result[i] = content[pivot: pivot + v];
  pivot += v;
 } 
return result;

}

func IntToString(num int) string{
	return strconv.Itoa(num)
}

func StringToInt(str string) int{
	num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Invalid number:", err)
        return 0
    }
    return num;
}
