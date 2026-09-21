func longestCommonPrefix(strs []string) string {
  prefix := "";
  baseString := strs[0];
  for i := range(len(baseString)){
	tempPrefix := prefix + string(baseString[i]);
	for _,str := range(strs){
		if strings.Index(str, tempPrefix) == -1{
			return prefix;
		}
	}
	prefix = tempPrefix; 
  } 

  return prefix;
    
}
