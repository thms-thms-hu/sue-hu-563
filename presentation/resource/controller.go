package resource

import (
	"github.com/gin-gonic/gin"
	"sue-hu-563/service/resource"
)

func ApplyResource(c *gin.Context) {
	// www.plantuml.https://www.plantuml.com/plantuml/svgrepo-refresh.svgcom/plantuml/png/hP5FQzj04CNlyoaUFjcGK9USAcX92Frp2IrroNqhZvs5rQuOhPoc-CDNmwwALrO3eJdOh6RUlFdpk-wYaTWtZamTW-3MMVQHsY5eQziQ7_6Tkz1BpRV1HmdEiSocDom7MtEHsrnzCNxhM3wfkpmLoVkxIRB8QZVsyUElA6co85iMl5jFL4qpcAIfi-JjUasK8vSVPiQlyEthvaU5Dv9JTAxHoibc5MxQrZsDqULOVG7s1ol1D-pZ1NRMSRSYCYvYbE0z5jtLWa1JK-qplLA5pnp_1fSdxYk2sfv4rDayTeCrQc4J6UNzRnwxx4EOwLqeAgJ7Sbm1W0kXXNwZIfj1eqMPsanfLSZFd51BNcdfETZBip64OoyUTLBEDIp1JydamhG1-oq11B3h624-i9mUK2N7uv6IaFH9gYpO7Ph5BBdxknbo6lnVSeFpy-Iqzml9QL2hNXMS-ZqBJWKl1uUCuz_GpYyo3Qc93v4XTl-GXkbCmt1XRskOeR7b3Z_xY5qG66pjRiV2lgIqvBtkyBiNpbkCqq7DAGsgl-lXhs_S7m00

	var applyResourceDTO applyResourceDTO
	if err := c.Bind(&applyResourceDTO); err != nil {
		// TODO: error handling
	}

	form, err := c.MultipartForm()
	if err != nil {
		// TODO: error handling
	}

	files := form.File["files"]
	if len(files) == 0 {
		// TODO: error handling
	}

	resource.ApplyResource(applyResourceDTO.Resource, applyResourceDTO.Environment, files)
}
