package lexer

// Lexer hold all relevent for curent token
type Lexer struct {
	input        string
	posistion    int  // current pos in input (points to current char)
	readPosition int  // current readin pos in input (after current char)
	ch           byte // current char under examination
}

// New returns a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.posistion = l.readPosition
	l.readPosition++
}
