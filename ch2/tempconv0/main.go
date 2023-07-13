package main

import "fmt"

type Celsius float64
type Fahrehit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrehit {
	return Fahrehit(c*9/5 + 32)

}

func FTtoC(f Fahrehit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Print(FTtoC(333.23))
}
