package resource

import (
	"github.com/gin-gonic/gin"
	"sue-hu-563/service/resource"
)

func ApplyResource(c *gin.Context) {
	// www.plantuml.com/plantuml/png/hL9DQnin5BphLmm-sH1Jbvowq98G-d49BNN9NLq_EmAjPDvgdQRuntVVhjQr59LYo251dZSpcgVPcpPe3brZbQw3PznPGoweVG-Qskor2_XEhU-ufZll0djhYOlZDV71r9JCrjTVjDlqYa-3-tCoVBZFXibKRYCNFluAF2Rpl276krN6QfhUP4mTAMzNAwKYilnGY5_XszVrZmflE8xKIuueiyqgtExtzdb2vvFw2kGEXhrho8KhR8sbTg6KjW4J1UynQwzd2YetbNloLwhmcS89d9-vBzGBCfQbf4oXO_UNh21npfmaaNvg-qYecNGWfBtTFhKdQI5pmgWm7kRJwW2iztl8RrGOO0YQoCGcRxb2FCR1OBaOf3aLopp6j7KzCaiNfXzCtcWKZXkGsoX8Qh8jmOT7ulC75ChnU5GZKSbH2JFf_D3CYeq_t1SQZs3UU4--eF4JMLrMUI_yN-N2kAXo0KNrceqdViN6XN1vuuWz_gFjuelaa9Wu7mXiTezZc5Ot1EjtfeRkBpRKucSNiFKCZOtPRed9fItDQISx_EwOuXPJki4yVd3oTzF_wnhx1m00
	// kubectl get all -l crossplane.io/managed-by-crossplane=true
	
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
