package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/app_error"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func BindData(ctx *gin.Context, req interface{}) bool {

	if ctx.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", ctx.FullPath())

		err := app_error.NewUnsupportMediaType(msg)

		ctx.JSON(err.Status(), gin.H{
			"error": err,
		},
		)
		return false
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.Printf("Could not able binding data: %v", err)

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []invalidArgument
			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			err := app_error.NewBadRequest("Invalid request parameters. Verify request parameters once.")

			ctx.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}
	}

	return true
}
