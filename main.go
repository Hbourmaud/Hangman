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
	letter := 65
	compteur := 0
	begin := 298
	cpt := 0
	scanner := bufio.NewScanner(file)
	if letter > 64 && letter < 91 {
		for i := 64; i < 91; i++ {
			if letter == i {
				break
			}
			compteur++
		}
		fmt.Println(compteur)
		begin = (begin + (8 * compteur)) + 1*compteur
		end := (begin + 8)
		//fmt.Println(begin)

		for scanner.Scan() {
			cpt++
			test = (scanner.Text())

			if cpt > begin {
				fmt.Println(test)
			}

			if cpt == end {
				os.Exit(0)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
