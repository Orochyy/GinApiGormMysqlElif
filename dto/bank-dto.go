package dto

type BankUpdateDTO struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Loan    string `json:"loan" form:"loan" binding:"required"`
	Percent string `json:"percent" form:"percent" binding:"required"`
	Term    string `json:"term" form:"term" binding:"required"`
}

type BankCreateDTO struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Loan    string `json:"loan" form:"loan" binding:"required"`
	Percent string `json:"percent" form:"percent" binding:"required"`
	Term    string `json:"term" form:"term" binding:"required"`
}