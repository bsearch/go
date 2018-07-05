package flipAndInvertImage

func FlipAndInvertImage(A [][]int) [][]int {
    for _,array := range A {
        isOdd:=len(array)%2!=0
        for i:= 0 ; i < len(array) / 2; i++ {
            array[i], array[len(array) - i - 1] = 1 - array[len(array) - i - 1], 1 - array[i]
        }
        if isOdd {
            array[len(array)/2] = 1 - array[len(array)/2]
        }
        
    }
    return A
}
