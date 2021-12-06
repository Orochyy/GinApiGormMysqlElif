package entity

type Bank struct {
	ID      uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name    string  `gorm:"type:varchar(255)" json:"name"`
	Loan    float64 `gorm:"type:float(33)" json:"loan"`
	Persent float64 `gorm:"type float(33)" json:"persent"`
	Term    float64 `gorm:"type:float(33)" json:"term"`
	UserID  uint64  `gorm:"not null" json:"-"`
	User    User    `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
