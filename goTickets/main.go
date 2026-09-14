package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"  gorm:"type:varchar(100);not null"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally return the error to let each route control the status code.
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func main() {
	dsn := "host=localhost user=postgres password=sujon123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("database connection successfully")
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/user", func(c *echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		result := db.Create(&u)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error.Error()})
		}
		return c.JSON(http.StatusCreated, u)
	})

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
