package lexer

type Lexer struct  {
	input string
	position int // curent position in input
	readPosition int // curent read position in input
	ch byte // current character under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}