// let five = 5;
// let ten = 10;
//
// let add = fn(x, y) {
//   x + y;
// };
//
// let result = add(five, ten);

package token

type TokenType string

type Token struct{
  Type    TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  //identifiers and literals
  IDENT   = "IDENT"
  INT     = "INT"

  //Operators
  ASSIGN  = "="
  PLUS    = "+"

  //Delimiters
  COMMA       = ","
  SEMICOLON   = ";"

  LPAREN      = "("
  RPAREN      = ")"

  LBRACE      = "{"
  RBRACE      = "}"

  //Keywords
  FUNCTION    = "FUNCTION"
  LET         = "LET"
)
