package api

import (
	"github.com/labstack/echo/v4"
	"github.com/roham96/go-basic/server/api/user"
	"github.com/roham96/go-basic/server/database"
)

func InitialBeans(e *echo.Echo, db *database.DB) {
	user.RegisterHandlers(e, user.NewService(user.NewRepository(db)))
}
