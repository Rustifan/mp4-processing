package parser

type Parser[T any] interface {
	Parser(data []byte) T
}
