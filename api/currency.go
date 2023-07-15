package api

import (
	"asia-yo-cur-svc/currency"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"strings"
)

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field(), e.Param())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}

type GetCurrencyRequest struct {
	Source string `form:"source" binding:"required"`
	Target string `form:"target" binding:"required"`
	Amount string `form:"amount" binding:"required"`
}

type GetCurrencyResponse struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

func (handler *Handler) GetCurrency(ctx *gin.Context) {
	var req GetCurrencyRequest
	var resp GetCurrencyResponse
	if err := ctx.ShouldBindQuery(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationError := validationErrors[0]
			resp.Msg = "failed, error: " + ValidationErrorToText(validationError)
			ctx.JSON(http.StatusBadRequest, resp)
			return
		} else {
			resp.Msg = "failed, error: Request body incorrect"
			ctx.JSON(http.StatusBadRequest, resp)
			return
		}
	}
	currencies, err := currency.LoadCurrencies("currencies.json")
	if err != nil {
		resp.Msg = "failed, error: load currency information"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := currencies.CheckCurrencyInUse(req.Source); err != nil {
		resp.Msg = "failed, error: " + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := currencies.CheckCurrencyInUse(req.Target); err != nil {
		resp.Msg = "failed, error: " + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	req.Amount = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(req.Amount, "")
	newAmount, err := currencies.Convertor(req.Source, req.Target, req.Amount)
	if err != nil {
		resp.Msg = "failed, error: " + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Msg = "success"
	resp.Amount = withComma(newAmount)
	ctx.JSON(http.StatusOK, resp)
}

func withComma(amount string) string {
	amountList := strings.Split(amount, ".")
	suffix, prefix := amountList[0], amountList[1]
	commaIdx := 3 - (len(suffix) % 3)
	if commaIdx == 3 {
		commaIdx = 0
	}
	var sb strings.Builder
	sb.WriteString("$")
	for i := 0; i < len(suffix); i++ {
		if commaIdx == 3 {
			sb.WriteString(",")
			commaIdx = 0
		}
		commaIdx++
		sb.WriteString(string(suffix[i]))
	}
	if len(prefix) > 0 {
		sb.WriteString(".")
		sb.WriteString(prefix)
	}
	return sb.String()
}
