package resource

import (
	"github.com/gin-gonic/gin"
	"sue-hu-563/service/resource"
)

func ApplyResource(c *gin.Context) {

	var applyResourceDTO applyResourceDTO
	if err := c.Bind(&applyResourceDTO); err != nil {
		// TODO: error handling
	}

	form, err := c.MultipartForm()
	if err != nil {
		// TODO: error handling
	}

	file := form.File["file"]
	if len(file) == 0 {
		// TODO: error handling
	}

	if len(file) != 1 {
		// TODO: error handling
	}

	resource.ApplyResource(applyResourceDTO.Resource, applyResourceDTO.Environment, file[0])
}
