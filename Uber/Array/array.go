package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
Problem :
Find the pairs across 2 sorted arrays with K largest sum

ex : // A={1, 2, 4, 5, 6}, B={3, 5, 7, 9} K = 3 result = (1, 3),(2, 3),(1, 5)
Algo: Since Array are sorted in acending order therefore there last elements sum would be the first Largest sum
Use index pointer on array to walk as we compare their pair, decrease index which pair
*/
