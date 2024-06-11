package middleware

import (
    "github.com/gofiber/fiber/v2"
    "log"
)

func (m *Middleware) ErrorHandler(ctx *fiber.Ctx, err error) error {
    // Log the error
    log.Printf("Error: %v", err)

    // Determine error type and send appropriate response
    code := fiber.StatusInternalServerError
    msg := "Internal Server Error"

    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
        msg = e.Message
    }

    return ctx.Status(code).JSON(fiber.Map{
        "error": msg,
    })
}
