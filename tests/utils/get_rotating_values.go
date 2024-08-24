package utils

import "fmt"

var e_idx int

func GetRotatingEmail() string {
	e_idx += 1
	return fmt.Sprintf("test_user%d@gmail.com", e_idx)
}

var ts_idx int
var taxShelters = []string{
	"TAXABLE",
	"ROTH",
	"TRADITIONAL",
	"HSA",
	"529",
}

func GetRotatingTaxShelter() string {
	ts_idx += 1
	return taxShelters[ts_idx%len(taxShelters)]
}

var inst_idx int
var institutions = []string{
	"Vanguard",
	"Fidelity",
	"Schwab",
	"Ameritrade",
}

func GetRotatingInstitution() string {
	inst_idx += 1
	return institutions[inst_idx%len(institutions)]
}

var dep_idx int

func GetRotatingDeprecation() bool {
	dep_idx += 1
	return dep_idx%10 == 0
}
