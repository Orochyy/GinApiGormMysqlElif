package controller

import (
	"GinApiGormMysqlElif/dto"
	"GinApiGormMysqlElif/entity"
	"GinApiGormMysqlElif/helper"
	"GinApiGormMysqlElif/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (c *bankController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var bank entity.Bank = c.bankService.FindByID(id)
	if (bank == entity.Bank{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", bank)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bankController) Insert(context *gin.Context) {
	var bankCreateDTO dto.BankCreateDTO
	errDTO := context.ShouldBind(&bankCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bankCreateDTO.UserID = convertedUserID
		}
		result := c.bankService.Insert(bankCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}
