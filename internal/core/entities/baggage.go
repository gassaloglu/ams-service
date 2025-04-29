package entities

type Baggage struct {
	ID               uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	BaggageID        string   `json:"baggage_id" db:"baggage_id"`
	BaggageAllowance *float64 `json:"baggage_allowance,omitempty" db:"baggage_allowance"`
	Weight           float64  `json:"weight" db:"weight"`
	Piece            int      `json:"piece" db:"piece"`
}
