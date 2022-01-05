package controller

import (
	"GinApiGormMysqlElif/entity"
	"GinApiGormMysqlElif/helper"
	"GinApiGormMysqlElif/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BankController interface {
}

type bankController struct {
	bankService service.BankService
	jwtService  service.JWTService
}

func NewBankController(bankServ service.BankService, jwtServ service.JWTService) BankController {
	return &bankController{
		bankService: bankServ,
		jwtService:  jwtServ,
	}
}

func (c *bankController) All(context *gin.Context) {
	var banks []entity.Bank = c.bankService.All()
	res := helper.BuildResponse(true, "OK", banks)
	context.JSON(http.StatusOK, res)
}
