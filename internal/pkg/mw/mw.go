package mw

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/logx/log"
)

func HandlerErrorLogging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if err := next(ctx); err != nil {
				log.Error(fmt.Sprintf("could not complete handler (%v)", next), err)

				return err
			}

			return nil
		}
	}
}
