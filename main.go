package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	word := (WordChoose())
	fmt.Println(string(word))
	fmt.Println("Good Luck, you have 10 attempts.")
	attempts := 10
	lword := len(word)
	tableau := []byte{}
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
					break
				}
			}
		}
	}
	PrintTable(tableau)
	fmt.Println()
	for {
		lettertest := EnterLetter()
		tableau, attempts = Check(tableau, word, lettertest, attempts)
		if CheckFin(tableau) {
			return
		}
		if attempts == 0 {
			return
		}
	}

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

func EnterLetter() byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose: ")
	letter, _ := reader.ReadString('\n')
	return byte(letter[0])
}

func Check(tableauV []byte, word []byte, letter byte, attempts int) ([]byte, int) {
	pres := false
	for i := 0; i < len(word); i++ {
		if letter == word[i] {
			tableauV[i] = letter
			pres = true
		}
	}
	if !pres {
		attempts--
		fmt.Print("Not present in the word, ")
		fmt.Print(attempts)
		fmt.Println(" attempts remaining")
		PrintHang(attempts)
	} else {
		PrintTable(tableauV)
		fmt.Println()
	}
	return tableauV, attempts
}

func CheckFin(tableau []byte) bool {
	for i := 0; i < len(tableau); i++ {
		if tableau[i] == '_' {
			return false
		}
	}
	fmt.Println("Congrats !")
	return true
}

func PrintTable(tableau []byte) {
	for i := 0; i < len(tableau); i++ {
		fmt.Print(string(tableau[i]))
		fmt.Print(string(' '))
	}
	fmt.Println()
}

func PrintHang(attempts int) {
	hang, _ := os.Open("hangman.txt")
	info, _ := os.Stat("hangman.txt")
	size := info.Size()
	arr := make([]byte, size)
	hang.Read(arr)
	hang.Close()
	nb := (attempts - 10) * -1
	fmt.Println(string(arr[0+80*(nb-1)-(1*(nb-1)) : 77+80*(nb-1)-(1*(nb-1))]))
}
