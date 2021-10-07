package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	word := (WordChoose())
	maj := false
	min := false
	//cheat code
	fmt.Println(string((word)))
	fmt.Println("Good Luck, you have 10 attempts.")
	attempts := 10
	lword := len(word)
	tableau := []byte{}
	for i := 0; i < lword; i++ {
		tableau = append(tableau, '_')
	}
	var letter byte
	var stockLetter []byte
	var compteur1 int
	if (word[0] < 91 && word[0] > 64) || (word[0] < 215 && word[0] > 191) || (word[0] < 221 && word[0] > 216) {
		maj = true
	} else if (word[0] < 123 && word[0] > 96) || (word[0] > 223 && word[0] < 247) || (word[0] > 248 && word[0] <= 255) {
		min = true
	}
	for i := 0; i < (len(word)/2)-1; i++ {
		compteur1 = 0
		letter = LetterAlea(word)
		for j := 0; j < len(stockLetter); j++ {
			if letter == stockLetter[j] {
				i--
				compteur1++
				break
			} else {
				continue
			}
		}
		if compteur1 > 0 {
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
		if min {
			// accent de 'e'
			if lettertest == 'e' || lettertest == 'è' || lettertest == 'é' || lettertest == 'ê' || letter == 'ë' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'e', compteur)
				tableau, compteur = Check(tableau, word, 'è', compteur)
				tableau, compteur = Check(tableau, word, 'é', compteur)
				tableau, compteur = Check(tableau, word, 'ê', compteur)
				tableau, compteur = Check(tableau, word, 'ë', compteur)
				fmt.Println(compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'E' en minuscune
			if lettertest == 'E' || lettertest == 'È' || lettertest == 'É' || lettertest == 'Ê' || letter == 'Ë' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'e', compteur)
				tableau, compteur = Check(tableau, word, 'è', compteur)
				tableau, compteur = Check(tableau, word, 'é', compteur)
				tableau, compteur = Check(tableau, word, 'ê', compteur)
				tableau, compteur = Check(tableau, word, 'ë', compteur)
				fmt.Println(compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'a'
			if lettertest == 'a' || lettertest == 'à' || lettertest == 'á' || lettertest == 'â' || letter == 'ã' || letter == 'ä' || letter == 'å' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'a', compteur)
				tableau, compteur = Check(tableau, word, 'à', compteur)
				tableau, compteur = Check(tableau, word, 'á', compteur)
				tableau, compteur = Check(tableau, word, 'â', compteur)
				tableau, compteur = Check(tableau, word, 'ã', compteur)
				tableau, compteur = Check(tableau, word, 'ä', compteur)
				tableau, compteur = Check(tableau, word, 'å', compteur)
				if compteur == 7 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'A' en minuscule
			if lettertest == 'A' || lettertest == 'À' || lettertest == 'Á' || lettertest == 'Â' || letter == 'Ã' || letter == 'Ä' || letter == 'Å' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'a', compteur)
				tableau, compteur = Check(tableau, word, 'à', compteur)
				tableau, compteur = Check(tableau, word, 'á', compteur)
				tableau, compteur = Check(tableau, word, 'â', compteur)
				tableau, compteur = Check(tableau, word, 'ã', compteur)
				tableau, compteur = Check(tableau, word, 'ä', compteur)
				tableau, compteur = Check(tableau, word, 'å', compteur)
				if compteur == 7 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'i'
			if lettertest == 'i' || lettertest == 'ì' || lettertest == 'í' || lettertest == 'î' || letter == 'ï' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'i', compteur)
				tableau, compteur = Check(tableau, word, 'ì', compteur)
				tableau, compteur = Check(tableau, word, 'í', compteur)
				tableau, compteur = Check(tableau, word, 'î', compteur)
				tableau, compteur = Check(tableau, word, 'ï', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'I' en minuscule
			if lettertest == 'I' || lettertest == 'Ì' || lettertest == 'Í' || lettertest == 'Î' || letter == 'Ï' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'i', compteur)
				tableau, compteur = Check(tableau, word, 'ì', compteur)
				tableau, compteur = Check(tableau, word, 'í', compteur)
				tableau, compteur = Check(tableau, word, 'î', compteur)
				tableau, compteur = Check(tableau, word, 'ï', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'o'
			if lettertest == 'o' || lettertest == 'ò' || lettertest == 'ó' || lettertest == 'ô' || letter == 'õ' || letter == 'ö' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'o', compteur)
				tableau, compteur = Check(tableau, word, 'ò', compteur)
				tableau, compteur = Check(tableau, word, 'ó', compteur)
				tableau, compteur = Check(tableau, word, 'ô', compteur)
				tableau, compteur = Check(tableau, word, 'õ', compteur)
				tableau, compteur = Check(tableau, word, 'ö', compteur)
				if compteur == 6 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'O' en minuscule
			if lettertest == 'O' || lettertest == 'Ò' || lettertest == 'Ó' || lettertest == 'Ô' || letter == 'Õ' || letter == 'Ö' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'o', compteur)
				tableau, compteur = Check(tableau, word, 'ò', compteur)
				tableau, compteur = Check(tableau, word, 'ó', compteur)
				tableau, compteur = Check(tableau, word, 'ô', compteur)
				tableau, compteur = Check(tableau, word, 'õ', compteur)
				tableau, compteur = Check(tableau, word, 'ö', compteur)
				if compteur == 6 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'u'
			if lettertest == 'u' || lettertest == 'ù' || lettertest == 'ú' || lettertest == 'û' || letter == 'ü' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'u', compteur)
				tableau, compteur = Check(tableau, word, 'ù', compteur)
				tableau, compteur = Check(tableau, word, 'ú', compteur)
				tableau, compteur = Check(tableau, word, 'û', compteur)
				tableau, compteur = Check(tableau, word, 'ü', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'U' en minuscule
			if lettertest == 'U' || lettertest == 'Ù' || lettertest == 'Ú' || lettertest == 'Û' || letter == 'Ü' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'u', compteur)
				tableau, compteur = Check(tableau, word, 'ù', compteur)
				tableau, compteur = Check(tableau, word, 'ú', compteur)
				tableau, compteur = Check(tableau, word, 'û', compteur)
				tableau, compteur = Check(tableau, word, 'ü', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
		} else {
			// accent de 'E'
			if lettertest == 'E' || lettertest == 'È' || lettertest == 'É' || lettertest == 'Ê' || letter == 'Ë' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'E', compteur)
				tableau, compteur = Check(tableau, word, 'È', compteur)
				tableau, compteur = Check(tableau, word, 'É', compteur)
				tableau, compteur = Check(tableau, word, 'Ê', compteur)
				tableau, compteur = Check(tableau, word, 'Ë', compteur)
				fmt.Println(compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'e' en majuscule
			if lettertest == 'e' || lettertest == 'è' || lettertest == 'é' || lettertest == 'ê' || letter == 'ë' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'E', compteur)
				tableau, compteur = Check(tableau, word, 'È', compteur)
				tableau, compteur = Check(tableau, word, 'É', compteur)
				tableau, compteur = Check(tableau, word, 'Ê', compteur)
				tableau, compteur = Check(tableau, word, 'Ë', compteur)
				fmt.Println(compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'A'
			if lettertest == 'A' || lettertest == 'À' || lettertest == 'Á' || lettertest == 'Â' || letter == 'Ã' || letter == 'Ä' || letter == 'Å' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'A', compteur)
				tableau, compteur = Check(tableau, word, 'À', compteur)
				tableau, compteur = Check(tableau, word, 'Á', compteur)
				tableau, compteur = Check(tableau, word, 'Â', compteur)
				tableau, compteur = Check(tableau, word, 'Ã', compteur)
				tableau, compteur = Check(tableau, word, 'Ä', compteur)
				tableau, compteur = Check(tableau, word, 'Å', compteur)
				if compteur == 7 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'a' en majuscule
			if lettertest == 'a' || lettertest == 'à' || lettertest == 'á' || lettertest == 'â' || letter == 'ã' || letter == 'ä' || letter == 'å' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'A', compteur)
				tableau, compteur = Check(tableau, word, 'À', compteur)
				tableau, compteur = Check(tableau, word, 'Á', compteur)
				tableau, compteur = Check(tableau, word, 'Â', compteur)
				tableau, compteur = Check(tableau, word, 'Ã', compteur)
				tableau, compteur = Check(tableau, word, 'Ä', compteur)
				tableau, compteur = Check(tableau, word, 'Å', compteur)
				if compteur == 7 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'I'
			if lettertest == 'I' || lettertest == 'Ì' || lettertest == 'Í' || lettertest == 'Î' || letter == 'Ï' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'I', compteur)
				tableau, compteur = Check(tableau, word, 'Ì', compteur)
				tableau, compteur = Check(tableau, word, 'Í', compteur)
				tableau, compteur = Check(tableau, word, 'Î', compteur)
				tableau, compteur = Check(tableau, word, 'Ï', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'i' en majuscule
			if lettertest == 'i' || lettertest == 'ì' || lettertest == 'í' || lettertest == 'î' || letter == 'ï' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'I', compteur)
				tableau, compteur = Check(tableau, word, 'Ì', compteur)
				tableau, compteur = Check(tableau, word, 'Í', compteur)
				tableau, compteur = Check(tableau, word, 'Î', compteur)
				tableau, compteur = Check(tableau, word, 'Ï', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'O'
			if lettertest == 'O' || lettertest == 'Ò' || lettertest == 'Ó' || lettertest == 'Ô' || letter == 'Õ' || letter == 'Ö' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'O', compteur)
				tableau, compteur = Check(tableau, word, 'Ò', compteur)
				tableau, compteur = Check(tableau, word, 'Ó', compteur)
				tableau, compteur = Check(tableau, word, 'Ô', compteur)
				tableau, compteur = Check(tableau, word, 'Õ', compteur)
				tableau, compteur = Check(tableau, word, 'Ö', compteur)
				if compteur == 6 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'o' en majuscule
			if lettertest == 'o' || lettertest == 'ò' || lettertest == 'ó' || lettertest == 'ô' || letter == 'õ' || letter == 'ö' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'O', compteur)
				tableau, compteur = Check(tableau, word, 'Ò', compteur)
				tableau, compteur = Check(tableau, word, 'Ó', compteur)
				tableau, compteur = Check(tableau, word, 'Ô', compteur)
				tableau, compteur = Check(tableau, word, 'Õ', compteur)
				tableau, compteur = Check(tableau, word, 'Ö', compteur)
				if compteur == 6 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'U'
			if lettertest == 'U' || lettertest == 'Ù' || lettertest == 'Ú' || lettertest == 'Û' || letter == 'Ü' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'U', compteur)
				tableau, compteur = Check(tableau, word, 'Ù', compteur)
				tableau, compteur = Check(tableau, word, 'Ú', compteur)
				tableau, compteur = Check(tableau, word, 'Û', compteur)
				tableau, compteur = Check(tableau, word, 'Ü', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
			// accent de 'u' en majuscule
			if lettertest == 'u' || lettertest == 'ù' || lettertest == 'ú' || lettertest == 'û' || letter == 'ü' {
				compteur := 0
				tableau, compteur = Check(tableau, word, 'U', compteur)
				tableau, compteur = Check(tableau, word, 'Ù', compteur)
				tableau, compteur = Check(tableau, word, 'Ú', compteur)
				tableau, compteur = Check(tableau, word, 'Û', compteur)
				tableau, compteur = Check(tableau, word, 'Ü', compteur)
				if compteur == 5 {
					attempts = abc(attempts)
					break
				}
				PrintTable(tableau)
			}
		}
		if lettertest != 'e' && lettertest != 'a' && lettertest != 'i' && lettertest != 'o' && lettertest != 'u' && (lettertest < 123 && lettertest > 96) || (lettertest < 91 && lettertest > 64) {
			compteur := -5
			tableau, compteur = Check(tableau, word, lettertest, compteur)
			if compteur == -1 {
				PrintTable(tableau)
			} else {
				attempts = abc(attempts)
			}
		}
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
	listword, _ := ioutil.ReadFile(args)
	info, _ := os.Stat(args)
	size := info.Size()
	arr := make([]byte, size)
	arr = listword
	var res []byte
	for i := 0; i < len(arr); i++ {
		if arr[i] == 13 {
			cptmot++
		}
	}
	for i := 0; i < len(arr); i++ {
		res = append(res, byte(arr[i]))
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
				tempstr := string(word)
				var word []byte
				for _, wordabc := range tempstr {
					word = append(word, byte(wordabc))
				}
				return word
			}
			if cpt == n {
				word = res[index+2 : i]
				tempstr := string(word)
				var word []byte
				for _, wordabc := range tempstr {
					word = append(word, byte(wordabc))
				}
				return word
			}
			if cpt == n-1 && cptmot+1 == n {
				var temp []byte
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
	tempstr := string(letter)
	var rep []byte
	for _, wordabc := range tempstr {
		rep = append(rep, byte(wordabc))
	}
	return rep[0]
}

func Check(tableauV []byte, word []byte, letter byte, compteur int) ([]byte, int) {
	pres := false
	for i := 0; i < len(word); i++ {
		if letter == word[i] {
			compteur = -1
			tableauV[i] = letter
			pres = true
		}
	}
	if !pres {
		compteur += 1
	}
	return tableauV, compteur
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

func abc(attempts int) int {
	attempts--
	fmt.Print("Not present in the word, ")
	fmt.Print(attempts)
	fmt.Println(" attempts remaining")
	PrintHang(attempts)
	return attempts
}
