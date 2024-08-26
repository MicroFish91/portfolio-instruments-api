package utils

import "fmt"

var e_idx int

func GetRotatingEmail() string {
	e_idx += 1
	return fmt.Sprintf("test_user%d@gmail.com", e_idx)
}

var taxShelters = []string{
	"TAXABLE",
	"ROTH",
	"TRADITIONAL",
	"HSA",
	"529",
}

func GetRotatingTaxShelter(ts_idx *int) string {
	*ts_idx += 1
	return taxShelters[*ts_idx%len(taxShelters)]
}

var institutions = []string{
	"Vanguard",
	"Fidelity",
	"Schwab",
	"Ameritrade",
}

func GetRotatingInstitution(inst_idx *int) string {
	*inst_idx += 1
	return institutions[*inst_idx%len(institutions)]
}

func GetRotatingDeprecation(dep_idx *int) bool {
	*dep_idx += 1
	return *dep_idx%10 == 0
}
