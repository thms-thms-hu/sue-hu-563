package resource

import (
	"github.com/gin-gonic/gin"
	"sue-hu-563/service/resource"
)

func ApplyResource(c *gin.Context) {
	// TODO: maybe use dto parsing?

	res := c.Request.FormValue("resource")
	env := c.Request.FormValue("environment")
	file, err := c.FormFile("file")

	if err != nil {
		// TODO: error handling
	}

	resource.ApplyResource(res, env, file)
}
