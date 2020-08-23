package main

import (
	"fmt"
)

type Stack struct{
	code [30000]rune
	mutableArray [1024]byte
	position int
}


func SyntaxAnalyzer(s *Stack){
	for i := 0; i < len(s.code); i++{
		if s.code[i] != 0 {
			switch char := s.code[i]; char{
				case '^':
					if s.mutableArray[s.position] >= 127 {
						s.mutableArray[s.position] = 0
					}else{
						s.mutableArray[s.position]++
						//fmt.Printf("Increase value at index: %d Value now: %d \n", s.position, s.mutableArray[s.position])
					}
				case 'v':
					if s.mutableArray[s.position] <= 0 {
						s.mutableArray[s.position] = 127
					}else{
						s.mutableArray[s.position]--
					}
				case '>':
					if s.position >= 0 {
						s.position++
					}
				case '<':
					if s.position <= len(s.mutableArray) {
						s.position--
					}
				case '?':
					fmt.Printf("%c", s.mutableArray[s.position])
				}
			}else{ break }
	}
}

/*func CodeRuner (s *Stack){
	for i := 0; i < len(s.mutableArray); i++{
		fmt.Printf("%c", s.mutableArray[i])
	}
}*/

func CodePusher(s *Stack){
	for i := 0; i < len(s.code); i++{
		fmt.Scanf("%c", &s.code[i])
		if s.code[i] == 'q'{
			s.code[i] = 0
			break
		}
	}
}

func main() {

	s := Stack{}

	CodePusher(&s)
	SyntaxAnalyzer(&s)
	//CodeRuner(&s)
}