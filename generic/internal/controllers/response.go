package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// STATUS OK
func Ok(ctx *fiber.Ctx, Data interface{}) error {
	ctx.Status(http.StatusOK).JSON(new(Data, "", true))
	return nil
}

// STATUS OK
func OkWithMessage(ctx *fiber.Ctx, Data interface{}, message string) error {
	ctx.Status(http.StatusOK).JSON(new(Data, message, true))
	return nil
}

// INTERNAL SERVER ERROR
func ServerError(ctx *fiber.Ctx, err error) error {
	ctx.Status(http.StatusInternalServerError).JSON(new(nil, err.Error(), true))
	return nil
}

// CUSTOM ERROR
func Error(ctx *fiber.Ctx, Status int, message string) error {
	ctx.Status(Status).JSON(new(nil, message, true))
	return nil
}

// CUSTOM ERROR
func NotFound(ctx *fiber.Ctx, err error) error {
	ctx.Status(http.StatusNotExtended).JSON(new(nil, err.Error(), true))
	return nil
}

// BAD REQUEST
func BadRequest(ctx *fiber.Ctx, err error) error {
	ctx.Status(http.StatusBadRequest).JSON(new(nil, err.Error(), true))
	return nil
}

// NEW RESPONSE
func new(Data interface{}, message string, Success bool) *Response {
	if Success && message == "" {
		message = "completed successfully"
	}
	return &Response{
		Message: message,
		Data:    Data,
		Success: Success,
	}
}
