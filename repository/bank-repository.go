package repository

import (
	"GinApiGormMysqlElif/entity"
	"gorm.io/gorm"
)

type BankRepository interface {
}

type bankConnection struct {
	connection *gorm.DB
}

func NewBankRepository(dbConn *gorm.DB) BankRepository {
	return &bankConnection{
		connection: dbConn,
	}
}

func (db *bankConnection) InsertBank(b entity.Bank) entity.Bank {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bankConnection) UpdateBank(b entity.Bank) entity.Bank {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bankConnection) DeleteBank(b entity.Bank) {
	db.connection.Delete(&b)
}

func (db *bankConnection) FindBankByID(bankID uint64) entity.Bank {
	var bank entity.Bank
	db.connection.Preload("User").Find(&bank, bankID)
	return bank
}

func (db *bankConnection) AllBank() []entity.Bank {
	var banks []entity.Bank
	db.connection.Preload("User").Find(&banks)
	return banks
}
