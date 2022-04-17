package ledger

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var symbols = map[string]string{
	"EUR": "€",
	"USD": "$",
}

type Locale struct {
	headerColumns      [3]string
	timeLayout         string
	thousandSeparator  string
	decimalSeparator   string
	positiveFormatting string
	negativeFormatting string
}

var locales = map[string]Locale{
	"en-US": {
		headerColumns:      [3]string{"Date", "Description", "Change"},
		timeLayout:         "02/01/2006",
		thousandSeparator:  ",",
		decimalSeparator:   ".",
		positiveFormatting: "%s%s ",
		negativeFormatting: "(%s%s)",
	},
	"nl-NL": {
		headerColumns:      [3]string{"Datum", "Omschrijving", "Verandering"},
		timeLayout:         "01-02-2006",
		thousandSeparator:  ".",
		decimalSeparator:   ",",
		positiveFormatting: "%s %s ",
		negativeFormatting: "%s %s-",
	},
}

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func (e *Entry) getDescription() string {
	if len(e.Description) > 25 {
		return fmt.Sprint(e.Description[:22], "...")
	}
	return e.Description
}

func (e *Entry) getChange(locale, currency string) string {
	symbol := symbols[currency]
	l := locales[locale]
	change := FormatMoney(e.Change, l.thousandSeparator, l.decimalSeparator)
	if e.Change >= 0 {
		return fmt.Sprintf(l.positiveFormatting, symbol, change)
	} else {
		return fmt.Sprintf(l.negativeFormatting, symbol, change)
	}
}

func FormatMoney(change int, thousandSeparator, decimalSeparator string) string {
	if change < 0 {
		change *= -1
	}

	number := change / 100
	decimals := change - number*100
	stringNumber := []rune(strconv.Itoa(number))
	formattedNumber := make([]rune, 0)

	for i := len(stringNumber) - 1; i >= 0; i-- {
		if len(formattedNumber) > 0 && len(formattedNumber)%3 == 0 {
			formattedNumber = append(formattedNumber, rune(thousandSeparator[0]))
		}
		formattedNumber = append(formattedNumber, stringNumber[i])
	}

	// reverse formattedNumber
	for i, j := 0, len(formattedNumber)-1; i < j; i, j = i+1, j-1 {
		formattedNumber[i], formattedNumber[j] = formattedNumber[j], formattedNumber[i]
	}

	return fmt.Sprint(string(formattedNumber), decimalSeparator, fmt.Sprintf("%02d", decimals))
}

func getHeader(locale string) string {
	l := locales[locale]
	date, description, change := l.headerColumns[0], l.headerColumns[1], l.headerColumns[2]
	return fmt.Sprintf("%-10s | %-25s | %s\n", date, description, change)
}

func getRow(currency, locale string, entry Entry) (string, error) {
	var date, description, change string
	t, err := time.Parse("2006-02-01", entry.Date)
	if err != nil {
		return "", err
	}
	l := locales[locale]
	change, date, description = entry.getChange(locale, currency), t.Format(l.timeLayout), entry.getDescription()
	return fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, change), nil
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "USD" && currency != "EUR" {
		return "", fmt.Errorf("invalid currency: %s", currency)
	}
	if locale != "en-US" && locale != "nl-NL" {
		return "", fmt.Errorf("invalid locale: %s", locale)
	}

	var entriesCopy []Entry
	for _, entry := range entries {
		entriesCopy = append(entriesCopy, entry)
	}
	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date == entriesCopy[j].Date {
			if entriesCopy[i].Description == entriesCopy[j].Description {
				return entriesCopy[i].Change < entriesCopy[j].Change
			}
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Date < entriesCopy[j].Date
	})

	var output strings.Builder
	output.WriteString(getHeader(locale))

	for _, entry := range entriesCopy {
		row, err := getRow(currency, locale, entry)
		if err != nil {
			return "", err
		}
		output.WriteString(row)
	}
	return output.String(), nil
}
