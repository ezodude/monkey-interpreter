package lexer

import(
  "github.com/ezodude/monkey/token"
)

type Lexer struct{
  input         string
  position      int
  readPosition  int
  ch            byte
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
}

func (l *Lexer) NextToken()  token.Token{
  return newToken("", 0)
}

func (l *Lexer) readChar()  {
  if l.readPosition >= len(l.input){
    l.ch = 0
  }else{
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte)  token.Token{
  return token.Token{Type: tokenType, Literal: string(ch)}
}
