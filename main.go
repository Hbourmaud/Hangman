package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println(WordChoose())
}

func WordChoose() []byte {
	var n int
	cpt := 0
	cptmot := 0
	index := 0
	args := os.Args[1]
	var word []byte
	listword, _ := os.Open(args)
	info, _ := os.Stat(args)
	size := info.Size()
	arr := make([]byte, size)
	listword.Read(arr)
	listword.Close()
	for i := 0; i < len(arr); i++ {
		if arr[i] == 13 {
			cptmot++
		}
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n = r1.Intn(cptmot + 1)
	if n == 0 {
		n++
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 13 {
			cpt++
			if n == 1 {
				word = arr[0:i]
				return word
			}
			if cpt == n {
				word = arr[index+2 : i]
				return word
			}
			if cpt == n-1 && cptmot+1 == n {
				var temp []byte
				word = arr[index+2:]
				for k := 0; k < len(word); k++ {
					if word[k] == 10 {
						temp = word[k+1:]
					}
				}
				return temp
			}
			index = i
		}
	}
	return word
}
