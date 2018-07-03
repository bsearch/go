func uniqueMorseRepresentations(words []string) int {
    rep := [] string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    
    bMap := make(map[string]int)
    for _,word := range words {
        var value string
        for _,char := range word {
            index := char - 'a';
            value += rep[index]
        }
        count := bMap[value]
        bMap[value] = count + 1
    }
    
    return len(bMap)
}
