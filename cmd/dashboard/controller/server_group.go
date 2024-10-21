package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/naiba/nezha/model"
	"github.com/naiba/nezha/service/singleton"
)

// List server group
// @Summary List server group
// @Schemes
// @Description List server group
// @Security BearerAuth
// @Tags common
// @Produce json
// @Success 200 {object} model.CommonResponse[[]model.ServerGroup]
// @Router /server-group [get]
func listServerGroup(c *gin.Context) error {
	authorizedUser, has := c.Get(model.CtxKeyAuthorizedUser)
	log.Println("bingo test", authorizedUser, has)
	var sg []model.ServerGroup
	if err := singleton.DB.Find(&sg).Error; err != nil {
		return err
	}

	c.JSON(http.StatusOK, model.CommonResponse[[]model.ServerGroup]{
		Success: true,
		Data:    sg,
	})
	return nil
}