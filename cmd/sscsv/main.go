package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/rmanzoku/sscsv"
)

func handler(ctx context.Context) (err error) {
	spreadsheetID := flag.Arg(0)
	sheet := flag.Arg(1)

	fmt.Println(spreadsheetID, sheet)

	c, err := sscsv.GetCSVFromSpreadsheet(spreadsheetID, sheet)
	if err != nil {
		return err
	}

	//	fmt.Println(c)

	f, err := os.Create(sheet + ".csv")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(c)

	return nil
}

func main() {
	flag.Parse()
	ctx := context.TODO()
	err := handler(ctx)
	if err != nil {
		panic(err)
	}
}
