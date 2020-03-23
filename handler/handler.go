package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
    return func(c echo.Context) error {
			jsonMap := map[string]string{
				"foo": "bar",
				"hoge": "fuga",
			}

			return c.JSON(http.StatusOK, jsonMap)
    }
}