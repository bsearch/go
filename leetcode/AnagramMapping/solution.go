func anagramMappings(A []int, B []int) []int {
    
    // create mapping of B key-> value (value->index)
    // return 
    
    var bMap map[int]int = make(map[int]int)
    for index, num := range B {
        bMap[num] = index
    }

    var res []int = make([] int, len(A))
    for index, num := range A {
        res[index] = bMap[num]
    }

    return res
}
