package entities

type Passenger struct {
	NationalId       string `json:"national_id" gorm:"size:11;not null"`
	PnrNo            string `json:"pnr_no" gorm:"size:6;not null"`
	BaggageAllowance int    `json:"baggage_allowance" gorm:"not null"`
	BaggageId        string `json:"baggage_id" gorm:"not null"`
	FareType         string `json:"fare_type" gorm:"not null"`
	Seat             string `json:"seat" gorm:"not null"`
	Meal             string `json:"meal" gorm:"not null"`
	ExtraBaggage     int    `json:"extra_baggage" gorm:"not null"`
	CheckIn          bool   `json:"check_in" gorm:"not null"`
	Name             string `json:"name" gorm:"size:50;not null"`
	Surname          string `json:"surname" gorm:"size:50;not null"`
	Email            string `json:"email" gorm:"size:100;not null"`
	Phone            string `json:"phone" gorm:"size:10;not null"`
	Gender           string `json:"gender" gorm:"not null"`
	BirthDate        string `json:"birth_date" gorm:"not null"`
	CipMember        bool   `json:"cip_member" gorm:"not null"`
	VipMember        bool   `json:"vip_member" gorm:"not null"`
	Disabled         bool   `json:"disabled" gorm:"not null"`
	Child            bool   `json:"child" gorm:"not null"`
}
