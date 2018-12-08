package tax

import (
	"regexp"
)

var rgxSQLInjectionChar = regexp.MustCompile(`[\'\"\t\r\0\n;]`) //to avoid sql injection

const (
	TaxFoodID          = 1
	TaxTobaccoID       = 2
	TaxEntertainmentID = 3
)
