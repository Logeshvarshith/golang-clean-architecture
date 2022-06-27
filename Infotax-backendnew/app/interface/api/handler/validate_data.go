package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

type InvalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// ValidateData function is used to validate the request headers and payload.
func ValidateData(ctx *gin.Context, req interface{}) bool {

	if ctx.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", ctx.FullPath())

		err := error.NewUnsupportMediaType(msg)

		ctx.JSON(err.Status(), gin.H{
			"error": err,
		},
		)
		return false
	}

	validate := validator.New()

	if err := ctx.ShouldBindJSON(req); err != nil {
		err := error.NewBadRequest("Invalid request parameters. Verify request parameters once.")
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return false
	}

	if err := validate.Struct(req); err != nil {
		log.Printf("Could not able binding data: %v", err)

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []InvalidArgument
			for _, err := range errs {
				invalidArgs = append(invalidArgs, InvalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			err := error.NewBadRequest("Invalid request parameters. Verify request parameters once.")

			ctx.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}
	}

	return true
}
