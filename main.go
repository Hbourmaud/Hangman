package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var test string
	letter := 68
	compteur := 0
	begin := 299
	cpt := 0
	scanner := bufio.NewScanner(file)
	if letter > 64 && letter < 91 {
		for i := 64; i < 91; i++ {
			if letter == i {
				break
			}
			compteur++
		}
		begin = (begin - 4) + (6 * compteur)
		//fmt.Println(begin)

		for scanner.Scan() {
			cpt++
			test = (scanner.Text())

			if cpt > begin {
				fmt.Println(test)
			}

			if cpt == (begin + 8) {
				os.Exit(0)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
