package resource

type applyResourceDTO struct {
	Environment string `form:"environment" binding:"required"`
	Resource    string `form:"resource" binding:"required"`
}
