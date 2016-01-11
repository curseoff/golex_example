// CAUTION: Generated file - DO NOT EDIT.

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	buf     []byte
	current byte
	src     *bufio.Reader
)

func getc() byte {
	buf = append(buf, current)
	current = 0
	if b, err := src.ReadByte(); err == nil {
		current = b
	}

	return current
}

func main() {
	fp, _ := os.Open(os.Args[1])
	defer fp.Close()

	src = bufio.NewReader(fp)
	c := getc()

yystate0:

	buf = buf[:0]

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = getc()
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= '-' || c == '/' || c >= ':' && c <= '@' || c >= '[' && c <= '`' || c >= '{' && c <= 'ÿ'
	case c == '.':
		goto yystate6
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate8
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate2:
	c = getc()
	goto yyrule5

yystate3:
	c = getc()
	goto yyrule6

yystate4:
	c = getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate5:
	c = getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate6:
	c = getc()
	switch {
	default:
		goto yyrule6
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate7:
	c = getc()
	switch {
	default:
		goto yyrule3
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate8:
	c = getc()
	switch {
	default:
		goto yyrule2
	case c == '.':
		goto yystate7
	case c >= '0' && c <= '9':
		goto yystate9
	}

yystate9:
	c = getc()
	switch {
	default:
		goto yyrule2
	case c == '.':
		goto yystate7
	case c >= '0' && c <= '9':
		goto yystate9
	}

yystate10:
	c = getc()
	switch {
	default:
		goto yyrule4
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate11
	}

yystate11:
	c = getc()
	switch {
	default:
		goto yyrule4
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate11
	}

yyrule1: // [ \t\n\r]+

	goto yystate0
yyrule2: // {D}
	{
		fmt.Printf("int %q\n", buf)
		goto yystate0
	}
yyrule3: // {D}\.{D}?|\.{D}
	{
		fmt.Printf("float %q\n", buf)
		goto yystate0
	}
yyrule4: // {A}
	{
		fmt.Printf("string %q\n", buf)
		goto yystate0
	}
yyrule5: // \0
	{
		return
	}
yyrule6: // .
	{
		fmt.Printf("unknown %q\n", buf)
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	panic("scanner internal error")
}
