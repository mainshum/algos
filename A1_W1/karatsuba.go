package A1_W1

import (
	"math"
	"math/big"
)

func getIntLen(num big.Int) int {
	return len(num.String())
}

func partitionString(str string, count int64) (string, string) {
	return str[0:count], str[count:len(str)]
}

func Karatsuba(x big.Int, y big.Int) big.Int {
	limit := big.NewInt(10)
	if x.Cmp(limit) < 0 && y.Cmp(limit) < 0 {
		return *x.Mul(&x, &y)
	}

	n := math.Max(float64(getIntLen(x)), float64(getIntLen(y)))
	m := int64(math.Ceil(n / 2.0))

	var exp, expTimesTwo, xH, xL, yH, yL big.Int
	exp.Exp(big.NewInt(10), big.NewInt(m), nil)
	expTimesTwo.Exp(big.NewInt(10), big.NewInt(m*2), nil)

	xH.DivMod(&x, &exp, &xL)
	yH.DivMod(&y, &exp, &yL)

	a := Karatsuba(xH, yH)
	d := Karatsuba(xL, yL)
	e := Karatsuba(*xH.Add(&xH, &xL), *yH.Add(&yH, &yL))

	e.Sub(&e, &d).Sub(&e, &a)

	firstFact := a.Mul(&a, &expTimesTwo)
	sndFact := e.Mul(&e, &exp)

	retval := firstFact.Add(firstFact, sndFact)
	retval.Add(retval, &d)

	return *retval
}
