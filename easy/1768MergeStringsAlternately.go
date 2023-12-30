func mergeAlternately(word1 string, word2 string) string {
    lenght1 := len(word1)
    lenght2 := len(word2)
    
    var mergedStr string
    var biggestNumber int
   
    if lenght1 >= lenght2 {
        biggestNumber = lenght1
    } else {
        biggestNumber = lenght2
    }

    for i := 0; i < biggestNumber; i ++ {
        if i < lenght1 {
            mergedStr += string(word1[i])
        }
        
        if i < lenght2 {
            mergedStr += string(word2[i])
        }
    } 
    
    return mergedStr
}
