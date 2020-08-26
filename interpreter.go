

1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
67
68
69
70
71
72
73
74
75
76
77
78
79
80
81
82
83
84
85
86
87
88
89
90
91
92
93
94
95
96
97
98
99
100
101
102
103
104
105
106
107

package main

import (
	"fmt"
	"unicode"
)

type stack struct{
	code [30000]rune
	mutableArray [1024]byte
	position int
}

func makeLoop(s *stack, start, end int) int {
	var skip = 0
	for i := start; i < end && s.code[i] != 0; i++ {
		switch s.code[i] {
		case '[':
			skip++
		case ']':
			if skip > 0 {
				skip--
				continue
			}

			for k := 0; k < 10; k++ {
				execCode(s, start, i)
			}

			return i
		}
	}

	panic("endless loop")
}

func execCode(s *stack, start, end int) {
	for i := start; i < end && s.code[i] != 0; i++ {
		switch s.code[i] {
		case '[':
			i = makeLoop(s, i+1, end)
		case '^':
			if s.mutableArray[s.position] >= 127 {
				s.mutableArray[s.position] = 0
				continue
			}
			s.mutableArray[s.position]++
			// fmt.Printf("Increase value at index: %d Value now: %d \n", s.position, s.mutableArray[s.position])
		case 'v':
				if s.mutableArray[s.position] <= 0 {
					s.mutableArray[s.position] = 127
					continue
				}
				s.mutableArray[s.position]--
		case '>':
			if s.position < len(s.mutableArray) - 1 {
				s.position++
			}
		case '<':
			if s.position > 0 {
				s.position--
			}
		case '?':
			fmt.Printf("%c\n", s.mutableArray[s.position])
		}
	}
}

func syntaxAnalyzer(s *stack) {
	execCode(s, 0, len(s.code))
}

/*func codeRuner (s *stack){
	for i := 0; i < len(s.mutableArray); i++{
		fmt.Printf("%c", s.mutableArray[i])
	}
}*/

func codePusher(s *stack){
	var symbol rune
	for i := 0; i < len(s.code); i++{
		fmt.Scanf("%c", &symbol)
		if unicode.IsSpace(symbol) {
			i--
			continue
		}
		if symbol == 'q'{
			break
		}
		s.code[i] = symbol
	}
}

func main() {

	s := &stack{}

	codePusher(s)
	syntaxAnalyzer(s)
	// codeRuner(s)

	// for _, v := range s.code{
	// 	if v == 0 {break}
	// 	fmt.Printf("%c\n", v)
	// }
}

