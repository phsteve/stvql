package main

import "bufio"
import "fmt"

import "os"

import "strings"
import (
	"github.com/davecgh/go-spew/spew"
	"github.com/xwb1989/sqlparser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("stvql> ")
		input, _ := reader.ReadString('\n')
		sql := strings.Trim(input, "\n")
		stmt, _ := sqlparser.Parse(sql)
		spew.Dump(stmt)
	}
}
