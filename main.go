package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	word := (WordChoose())
	fmt.Println("Good Luck, you have 10 attempts.")
	tableau := []byte{}
	lword := len(word)
	for i := 0; i < lword; i++ {
		tableau = append(tableau, '_')
	}

	var letter byte
	var stockLetter []byte
	var compteur int
	for i := 0; i < (len(word)/2)-1; i++ {
		compteur = 0
		letter = LetterAlea(word)
		for j := 0; j < len(stockLetter); j++ {
			if letter == stockLetter[j] {
				i--
				compteur++
				break
			} else {
				continue
			}
		}
		if compteur > 0 {
			continue
		} else {
			for j := 0; j < lword; j++ {
				if word[j] == letter {
					tableau[j] = letter
					stockLetter = append(stockLetter, letter)
					fmt.Println(letter)
					break
				}
			}
		}
	}

	fmt.Println(word)
	fmt.Println(string(tableau))
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

func LetterAlea(word []byte) byte {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(len(word))
	letter := word[n]
	return letter
}
