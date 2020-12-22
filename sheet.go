package googletablesfunction

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func getClient(ctx context.Context) (*http.Client, error) {
	return google.DefaultClient(ctx, "https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/spreadsheets.readonly")
}

func getSheetService(ctx context.Context) (*sheets.Service, error) {
	client, err := getClient(ctx)
	if err != nil {
		fmt.Printf("getClient error: %v", err.Error())
		return nil, err
	}
	return sheets.New(client)
}

func getWriteRange() string {
	return "A1:" + os.Getenv("RANGE_END_LETTER")
}

func getReadRange() string {
	return getWriteRange() + os.Getenv("RANGE_MAX_NUMBER")
}

func getSheet(ctx context.Context) ([]*Contact, error) {
	srv, err := getSheetService(ctx)
	if err != nil {
		fmt.Printf("sheets.New error: %v", err.Error())
		return nil, err
	}
	readRange := getReadRange()
	resp, err := srv.Spreadsheets.Values.Get(os.Getenv("SPREADSHEET_ID"), readRange).Context(ctx).Do()
	if err != nil {
		fmt.Printf("srv.Spreadsheets.Values.Get error: %v", err.Error())
		return nil, err
	}
	var contacts []*Contact
	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			contacts = append(contacts, &Contact{
				Email:     fmt.Sprintf("%v", row[0]),
				Telephone: fmt.Sprintf("%v", row[1]),
				Address:   fmt.Sprintf("%v", row[2]),
			})
		}
	}

	return contacts, nil
}

func writeSheet(ctx context.Context, contacts []*Contact) error {
	srv, err := getSheetService(ctx)
	if err != nil {
		fmt.Printf("sheets.New error: %v", err.Error())
		return err
	}
	rng := getWriteRange()
	rng += strconv.Itoa(len(contacts))
	values := make([][]interface{}, len(contacts))
	for i, c := range contacts {
		values[i] = []interface{}{c.Email, c.Telephone, c.Address}
	}
	valueInputOption := "RAW"
	vr := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Range:          rng,
		Values:         values,
	}
	_, err = srv.Spreadsheets.Values.Update(os.Getenv("SPREADSHEET_ID"), rng, vr).ValueInputOption(valueInputOption).Context(ctx).Do()
	if err != nil {
		fmt.Printf("srv.Spreadsheets.Values.Update error: %v", err.Error())
		return err
	}
	return nil
}
