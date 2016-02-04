
package main

import (
	"fmt"
  "sort"
)

func main (){

var numbers sort.Float64Slice = []float64{9,2,12,1,16,12,10}

numbers.Sort()  //numbers are sorted and can be used as is.

result := median(numbers)
fmt.Println(result)
}

func median(numbers []float64) float64 {
    middle := len(numbers) / 2
    result := numbers[middle]
    if len(numbers)%2 == 0 {
        result = (result + numbers[middle-1]) / 2
    }
    return result
}
