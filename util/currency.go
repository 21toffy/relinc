package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	NRA = "NRA"
)

func IsSupportedCurency(curency string) bool {
	switch curency {
	case USD, EUR, CAD, NRA:
		return true
	}
	return false
}
