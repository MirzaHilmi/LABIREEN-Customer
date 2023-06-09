package handlers

import (
	"labireen-customer/pkg/jwtx"
	"labireen-customer/pkg/response"
	"labireen-customer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	GetMe(ctx *gin.Context)
}

type customerHandlerImpl struct {
	svc services.CustomerService
}

func NewCustomerHandler(svc services.CustomerService) *customerHandlerImpl {
	return &customerHandlerImpl{svc}
}

func (cH *customerHandlerImpl) GetMe(ctx *gin.Context) {
	temp, exist := ctx.Get("currentUser")
	if !exist {
		log := response.ErrorLog{
			Field:   "token",
			Message: "Key token value does not exists",
		}
		response.Error(ctx, http.StatusNotFound, "Key error", log)
	}

	user := temp.(jwtx.CustomerClaims)

	userResp, err := cH.svc.GetCustomer(user.ID)
	if err != nil {
		log := response.ErrorLog{
			Field:   "ID",
			Message: "ID sent is not valid",
		}
		response.Error(ctx, http.StatusNotFound, "Cannot find requested data", log)
		return
	}

	response.Success(ctx, http.StatusOK, "success", userResp)
}
