package sscsv

import (
	"context"
	"errors"
	"os"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	filename = os.Getenv("HOME") + "/.config/cred.json"
)

func GetCSVFromSpreadsheet(spreadsheetID, sheet string) (csv string, err error) {

	srv, err := sheets.NewService(context.TODO(), option.WithCredentialsFile(filename))
	if err != nil {
		return
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, sheet).Do()
	if err != nil {
		return
	}

	if len(resp.Values) == 0 {
		return "", errors.New("No data found")
	}

	rows := make([]string, len(resp.Values))
	for i, row := range resp.Values {
		colomns := make([]string, len(row))
		for j, c := range row {
			s := Escape(c.(string))
			colomns[j] = s
		}
		rows[i] = strings.Join(colomns, ",")
	}

	csv = strings.Join(rows, "\n")
	return
}

func Escape(s string) string {
	x := strings.Replace(s, ",", "_+_", -1)
	x = strings.Replace(x, "\n", "\\n", -1)

	return x
}

func Unescape(s string) string {
	x := strings.Replace(s, "_+_", ",", -1)
	x = strings.Replace(x, "\\n", "\n", -1)

	return x
}
