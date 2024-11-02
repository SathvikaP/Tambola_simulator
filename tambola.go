// open github codespaces or any online ide and run go run tambola.go to see the results
package main

import (
	"fmt"
	"math/rand"
	"time"
	
)

type ticket [3][8]int


func main (){
	t:= ticket{}
	//idx1:= fiveUniqueIndices()
	//fmt.Println(idx1)
	uniqueVals:= make(map[int]bool)
	for row,_:= range t{
		idxs:= fiveUniqueIndices()
		vals:= fiveUniqueNums(uniqueVals)
		for i,val:= range idxs{
			t[row][val]= vals[i]
		}

	}
	fmt.Println(t)
	
}
func fiveUniqueIndices() [5]int{
	rand.Seed(time.Now().UnixNano()) 
	uniques:= make(map[int]bool)
	var idxs [5]int
	for i:=0;i<5; {
		
		n:= rand.Intn(8)
		if !uniques[n]{
			uniques[n]=true
			idxs[i]=n
			i++
		}
	}
	return idxs
}
func fiveUniqueNums(uniqueVals map[int]bool) [5]int{
	rand.Seed(time.Now().UnixNano())
	var values [5]int
	
	for i:=0;i<5; {
		n:= rand.Intn(100)+1
		if !uniqueVals[n]{
			uniqueVals[n]= true
			values[i]=n
			i++

		}
	}
	return values

}
