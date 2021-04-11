package handler

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/constant"
)

type response struct {
	Status  int                   `json:"status"`
	Code    constant.ResponseCode `json:"code"`
	Message string                `json:"message"`
	Data    interface{}           `json:"data"`
}

//设置返回内容
func setContent(c *fiber.Ctx, statusCode int, resp response) error{
	return c.Status(statusCode).JSON(&resp)
}

// SuccessWithMessage 请求成功时返回的内容
func SuccessWithMessage(c *fiber.Ctx, message string, data interface{}) error{
	resp := response{
		Status:  1,
		Code:    constant.SUCCESS,
		Message: message,
		Data:    data,
	}
	return setContent(c, fiber.StatusOK, resp)
}

// ErrorWithMessage 出现错误时返回的内容
func ErrorWithMessage(c *fiber.Ctx, responseCode constant.ResponseCode, message string) error{
	resp := response{
		Status:  -1,
		Code:    responseCode,
		Message: message,
		Data:    [0]int{},
	}
	return setContent(c, fiber.StatusOK, resp)
}

// StatusServerErrorWithMessage 服务器错误时返回的内容
func StatusServerErrorWithMessage(c *fiber.Ctx, message string) error{
	resp := response{
		Status:  -1,
		Code:    constant.CODE_500,
		Message: message,
		Data:    [0]int{},
	}
	return setContent(c, fiber.StatusInternalServerError, resp)
}
