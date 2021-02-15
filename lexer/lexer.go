package lexer

type Lexer struct {
	input string
	position int //current position in inut
	reaPosition int //current reading position in input
	ch byte //current char under examination
}

func New(input string) {
	l := &Lexer{input: input}
	return 1
}