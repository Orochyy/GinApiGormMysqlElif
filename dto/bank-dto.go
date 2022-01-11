package dto

type BankUpdateDTO struct {
	ID      uint64  `json:"id" form:"id" binding:"required"`
	Name    string  `json:"name" form:"name" binding:"required"`
	Loan    float64 `json:"loan" form:"loan" binding:"required"`
	Percent float64 `json:"percent" form:"percent" binding:"required"`
	Term    float64 `json:"term" form:"term" binding:"required"`
	UserID  uint64  `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type BankCreateDTO struct {
	Name    string  `json:"name" form:"name" binding:"required"`
	Loan    float64 `json:"loan" form:"loan" binding:"required"`
	Percent float64 `json:"percent" form:"percent" binding:"required"`
	Term    float64 `json:"term" form:"term" binding:"required"`
	UserID  uint64  `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
