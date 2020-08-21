package main


import (
	"fmt"
)



func main(){

		var bytes [1024]byte
		var char rune
		var position int = 0
		
		for {
			
			fmt.Scanf("%c", &char)
			
			if char == '^' && bytes[position] < 127{
				bytes[position]++
				fmt.Printf("Increase array value by 1 at addres: %p \n Value is now: %d \n Cell position: %d \n", &bytes[position], bytes[position], position)
			}
			
			if char == 'v' && bytes[position] > 0 {
				bytes[position]--
				fmt.Printf("Decrease array value by 1 at addres: %p \n Value is now: %d \n Cell position: %d \n", &bytes[position], bytes[position], position)
			}

			if char == '>' && position < len(bytes) {
				position++
				fmt.Printf("Move to next cell of array. \n Addres: %p \n Cell position: %d \n", &bytes[position], position)
			}

			if char == '<' && position > 0 {
				position--
				fmt.Printf("Move to previous cell of array. \n Addres: %p \n Cell position: %d \n", &bytes[position], position)
			}

			if char == '?' &&  bytes[position] >= 33 && bytes[position] < 127 {
				fmt.Printf("Char is: %c \n", bytes[position])
			}

			if char == '!' && bytes[position] >= 33 && bytes[position] < 127 {
				for i := 0; i < len(bytes); i++{
					if bytes[i] > 0 {
						fmt.Printf("%c", bytes[i])
					}else{ break; }
				}
				fmt.Printf("\n")
			}
		}

}