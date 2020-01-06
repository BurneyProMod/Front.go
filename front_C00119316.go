package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

var reader *bufio.Reader

const (
	LETTER  = 10
	DIGIT   = 1
	UNKNOWN = 99
	
	INT_LIT     = 10
	IDENT      = 11
	ADD_OP      = 21
	SUB_OP      = 22
	MULT_OP     = 23
	DIV_OP      = 24
	LEFT_PAREN  = 25
	RIGHT_PAREN = 26
	
	eof        = -1
)

var(
	nextToken int
	token int
	lexeme string
	nextChar rune = ' '
	charClass int

	file *os.File
	gerr error
)


func main() {
	var file, gerr = os.Open("/Desktop/Test/front.in")
	if gerr != nil {
		panic(gerr)
	}
	reader = bufio.NewReader(file)
	for {
		if gerr == io.EOF {
			fmt.Printf("Error: io.EOF")
			break
		} 
		else if gerr != nil {
			fmt.Printf("Error: gerr != nil")
			break
		}
		if lex() == eof {
			break
		}
	}
}

func lex() int {
	lexeme = ""
	getNonBlank()
	switch charClass {
	case LETTER:
		myaddChar()
		mygetChar()
		for charClass == LETTER || charClass == DIGIT {
			myaddChar()
			mygetChar()
		}
		nextToken = IDENT
	case DIGIT:
		myaddChar()
		mygetChar()
		for charClass == DIGIT {
			myaddChar()
			mygetChar()
		}
		nextToken = INT_LIT
	case UNKNOWN:
		lookup(nextChar)
		mygetChar()
	default:
		nextToken = eof
		lexeme = "EOF"
	}
	fmt.Println("Next token: ", nextToken, ",", "Next lexeme: ", lexeme)
	return nextToken
}

func myaddChar() {
	lexeme += string(nextChar)
}


func getNonBlank() {
	for unicode.IsSpace(nextChar) {
		mygetChar()
	}
}

func mygetChar() {
	ch, err := reader.ReadByte()
	gerr = err
	if err == nil || err == io.EOF {
		nextChar = rune(ch)
		if unicode.IsLETTER(nextChar) == true {
			charClass = LETTER
		} else if unicode.IsDIGIT(nextChar) == true {
			charClass = DIGIT
		} else if err == io.EOF {
			charClass = eof
		} else {
			charClass = UNKNOWN
		}
	} else {
		nextChar = '0'
		charClass = eof
	}
}

func lookup(ch rune) {
	switch ch {
	case '(':
		myaddChar()
		nextToken = LEFT_PAREN
	case ')':
		myaddChar()
		nextToken = RIGHT_PAREN
	case '+':
		myaddChar()
		nextToken = ADD_OP
	case '-':
		myaddChar()
		nextToken = SUB_OP
	case '*':
		myaddChar()
		nextToken = MULT_OP
	case '/':
		myaddChar()
		nextToken = DIV_OP
	default:
		myaddChar()
		nextToken = UNKNOWN
	}
}