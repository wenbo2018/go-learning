package main

import "math"
import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}

func main() {
	hmap := make(map[string]string);
	hmap["shen"] = "shenwenbo"
	fmt.Print(hmap["shen"])
}
