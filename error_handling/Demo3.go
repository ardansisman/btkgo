package error_handling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d --- %s", b.parameter, b.message)
}

func Demo3(tahmin int) (string, error) {
	if tahmin < 0 || tahmin > 100 {
		return "", &borderException{tahmin, "Sınırların dışında kaldı"}
	}

	return "Bildiniz", nil
}
