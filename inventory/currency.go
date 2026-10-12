package inventory

import "fmt"

type (
	Currency interface {
		Value() GoldPiece
	}
	CopperPiece   float32 // 1/100 GP
	SilverPiece   float32 // 1/10 GP
	ElectrumPiece float32 // 1/2 GP
	GoldPiece     float32 // 1 GP
	PlatinumPiece float32 // 10 GP
)

// Conversions

func (c CopperPiece) Value() GoldPiece {
	return GoldPiece(c / 100)
}

func (c SilverPiece) Value() GoldPiece {
	return GoldPiece(c / 10)
}

func (c ElectrumPiece) Value() GoldPiece {
	return GoldPiece(c / 2)
}

func (c GoldPiece) Value() GoldPiece {
	return c
}

func (c PlatinumPiece) Value() GoldPiece {
	return GoldPiece(c * 10)
}

// Strings

func (c CopperPiece) String() string {
	return fmt.Sprintf("%g CP", c)
}

func (c SilverPiece) String() string {
	return fmt.Sprintf("%g (SP)", c)
}

func (c ElectrumPiece) String() string {
	return fmt.Sprintf("%g (EP)", c)
}

func (c GoldPiece) String() string {
	return fmt.Sprintf("%g (GP)", c)
}

func (c PlatinumPiece) String() string {
	return fmt.Sprintf("%g (PP)", c)
}
