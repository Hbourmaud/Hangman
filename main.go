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
	fmt.Println(word)
	maj := false
	min := false
	fmt.Println(string((word)))
	fmt.Println("Good Luck, you have 10 attempts.")
	attempts := 10
	lword := len(word)
	tableau := []rune{}
	for i := 0; i < lword; i++ {
		tableau = append(tableau, '_')
	}
	var letter rune
	var stockLetter []rune
	var compteur int
	if word[0] < 91 && word[0] > 64 {
		maj = true
	} else if word[0] < 123 && word[0] > 96 {
		min = true
	}
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
		if maj && (lettertest < 123 && lettertest > 96) {
			lettertest = lettertest - 32
		}
		if min && (lettertest < 91 && lettertest > 64) {
			lettertest = lettertest + 32
		}
		tableau, attempts = Check(tableau, word, lettertest, attempts)
		if CheckFin(tableau) {
			return
		}
		if attempts == 0 {
			return
		}
	}

}

func WordChoose() []rune {
	var n int
	cpt := 0
	cptmot := 0
	index := 0
	args := os.Args[1]
	var word []rune
	listword, _ := os.Open(args)
	info, _ := os.Stat(args)
	size := info.Size()
	arr := make([]byte, size)
	listword.Read(arr)
	listword.Close()
	var res []rune
	for i := 0; i < len(arr); i++ {
		if arr[i] == 13 {
			cptmot++
		}
	}
	for i := 0; i < len(arr); i++ {
		res = append(res, rune(arr[i]))
	}
	if cptmot == 0 {
		return res
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n = r1.Intn(cptmot + 1)
	if n == 0 {
		n++
	}
	for i := 0; i < len(res); i++ {
		if arr[i] == 13 {
			cpt++
			if n == 1 {
				word = res[0:i]
				return word
			}
			if cpt == n {
				word = res[index+2 : i]
				return word
			}
			if cpt == n-1 && cptmot+1 == n {
				var temp []rune
				word = res[index+2:]
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

func LetterAlea(word []rune) rune {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(len(word))
	letter := word[n]
	return letter
}

func EnterLetter() rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose: ")
	letter, _ := reader.ReadString('\n')
	return rune(letter[0])
}

func Check(tableauV []rune, word []rune, letter rune, attempts int) ([]rune, int) {
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

func CheckFin(tableau []rune) bool {
	for i := 0; i < len(tableau); i++ {
		if tableau[i] == '_' {
			return false
		}
	}
	fmt.Println("Congrats !")
	return true
}

func PrintTable(tableau []rune) {
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
