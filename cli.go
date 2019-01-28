package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/xwb1989/sqlparser"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("stvql> ")
		input, _ := reader.ReadString('\n')
		raw := strings.Trim(input, "\n")
		sql, _ := sqlparser.Parse(raw)
		switch stmt := sql.(type) {
		case *sqlparser.DDL:
			// get the column names then initialize a CSV with them
			tableName := stmt.NewName.Name
			columnNames := make([]sqlparser.ColIdent, 0)
			for _, column := range stmt.TableSpec.Columns {
				columnNames = append(columnNames, column.Name)
			}
			f, _ := os.Create(fmt.Sprintf("./%s", tableName))
			defer f.Close()
			csvWriter := csv.NewWriter(f)
			defer csvWriter.Flush()
			toWrite := []string{}
			for _, value := range columnNames {
				toWrite = append(toWrite, sqlparser.String(value))
			}
			err := csvWriter.Write(toWrite)
			checkError("Cannot write to file", err)
			csvWriter.Flush()
		case *sqlparser.Select:
			fmt.Println("***got a select***")
		case *sqlparser.Insert:
			fmt.Println("***got an insert***")
		}
	}
}

func checkError(message string, err error) {
	if err != nil {
		fmt.Println("WTF")
	}
}
