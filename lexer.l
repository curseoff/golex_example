%{
package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    buf      []byte
    current  byte
    src *bufio.Reader
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
%}

%yyc c
%yyn c = getc()

D   [0-9]+
A   [a-zA-Z]+

%%
    buf = buf[:0] 

[ \t\n\r]+

{D}                 fmt.Printf("int %q\n", buf)

{D}\.{D}?|\.{D}     fmt.Printf("float %q\n", buf)

{A} fmt.Printf("string %q\n", buf)

\0                  return

.                   fmt.Printf("unknown %q\n", buf)

%%
    panic("scanner internal error")
}