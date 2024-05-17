package presentation

import (
	"github.com/gin-gonic/gin"
	"sue-hu-563/service"
)

func ApplyResource(c *gin.Context) {
	// TODO: maybe use dto parsing?

	resource := c.Request.Form.Get("resource")
	environment := c.Request.Form.Get("environment")
	file, err := c.FormFile("file")

	if err != nil {
		// TODO: error handling
	}

	// TODO: maybe multiple files
	service.ApplyResource(resource, environment, file)
}
