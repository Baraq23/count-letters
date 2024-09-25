package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countString("Heeeeeellloo TTTTherrree!!"))
	fmt.Println(countString("YoouuungFellllas"))
}

func countString(s string) string {
	newS := ""
	cnt := 1

	for i, c :=range s {
		if i == len(s)-1 {
			n := strconv.Itoa(cnt)
			newS +=n+string(c)
		}
		if i+1 != len(s) {
			if s[i+1] == byte(c){
				cnt++
			}else{
				n := strconv.Itoa(cnt)
				newS +=n+string(c)
				cnt = 1
			}
		}
	}
	return newS
}
