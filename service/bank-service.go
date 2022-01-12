package service

import (
	"GinApiGormMysqlElif/dto"
	"GinApiGormMysqlElif/entity"
	"GinApiGormMysqlElif/repository"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
	"math"
)

type BankService interface {
	Insert(b dto.BankCreateDTO) entity.Bank
	Update(b dto.BankUpdateDTO) entity.Bank
	Delete(b entity.Bank)
	All() []entity.Bank
	FindByID(bankID uint64) entity.Bank
	IsAllowedToEdit(userID string, bankID uint64) bool
	CountCreditPercents(bankID uint64) float64
}

type bankService struct {
	bankRepository repository.BankRepository
}

func NewBankService(bankRepo repository.BankRepository) BankService {
	return &bankService{
		bankRepository: bankRepo,
	}
}

func (service *bankService) Insert(b dto.BankCreateDTO) entity.Bank {
	bank := entity.Bank{}
	err := smapping.FillStruct(&bank, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bankRepository.InsertBank(bank)
	return res
}

func (service *bankService) Update(b dto.BankUpdateDTO) entity.Bank {
	bank := entity.Bank{}
	err := smapping.FillStruct(&bank, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	ress := service.bankRepository.UpdateBank(bank)
	return ress
}

func (service *bankService) Delete(b entity.Bank) {
	service.bankRepository.DeleteBank(b)
}

func (service *bankService) All() []entity.Bank {
	return service.bankRepository.AllBank()
}

func (service *bankService) FindByID(bankID uint64) entity.Bank {
	return service.bankRepository.FindBankByID(bankID)
}

func (service *bankService) IsAllowedToEdit(userID string, bankID uint64) bool {
	bank := service.bankRepository.FindBankByID(bankID)
	id := fmt.Sprintf("%v", bank.UserID)
	return userID == id
}

func (service *bankService) CountCreditPercents(bankID uint64) float64 {
	bank := service.bankRepository.FindBankByID(bankID)

	loan := bank.Loan
	percent := bank.Percent
	term := bank.Term

	result := (loan * (percent / 12)) * math.Pow(1+(percent/12), term) / (math.Pow(1+(percent/12), term) - 1)

	return result
}
