package entities

type Passenger struct {
	NationalId       string `json:"national_id" binding:"required,len=11,numeric"`
	PnrNo            string `json:"pnr_no" binding:"required,len=6,alphanum"`
	BaggageAllowance int    `json:"baggage_allowance" binding:"required,min=0,max=50"`
	BaggageId        string `json:"baggage_id" binding:"required,alphanum"`
	FareType         string `json:"fare_type" binding:"required,oneof=essentials advantage comfort"`
	Seat             string `json:"seat" binding:"required,alphanum"`
	Meal             string `json:"meal" binding:"required"`
	ExtraBaggage     int    `json:"extra_baggage" binding:"required"`
	CheckIn          bool   `json:"check_in" binding:"required"`
	Name             string `json:"name" binding:"required,alpha,min=2,max=50"`
	Surname          string `json:"surname" binding:"required,alpha,min=2,max=50"`
	Email            string `json:"email" binding:"required,email"`
	Phone            string `json:"phone" binding:"required,numeric,len=10"`
	Gender           string `json:"gender" binding:"required,oneof=male female other"`
	BirthDate        string `json:"birth_date" binding:"required,datetime=2006-01-02"`
	CipMember        bool   `json:"cip_member" binding:"required"`
	VipMember        bool   `json:"vip_member" binding:"required"`
	Disabled         bool   `json:"disabled" binding:"required"`
	Child            bool   `json:"child" binding:"required"`
}

type OnlineCheckInRequest struct {
	PNR     string `json:"pnr" binding:"len=6,alphanum"`
	Surname string `json:"surname" binding:"alpha,min=2,max=50"`
}

type GetPassengerByPnrRequest struct {
	PNR     string `form:"pnr" binding:"required,len=6,alphanum"`
	Surname string `form:"surname" binding:"required,alpha,min=2,max=50"`
}

type GetPassengerByIdRequest struct {
	NationalId string `form:"national_id" binding:"required,len=11,numeric"`
}
