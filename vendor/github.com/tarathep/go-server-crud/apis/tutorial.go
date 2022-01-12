package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/tarathep/go-server-crud/db"
	"github.com/tarathep/go-server-crud/model"
)

type TutorialHandler struct {
	DB db.TutorialRepository
}

func (h *TutorialHandler) CreateTutorial(c *gin.Context) {

	tutorial := model.Tutorial{}
	if err := c.ShouldBindJSON(&tutorial); err != nil {
		c.String(500, err.Error())
		return
	}

	if err := h.DB.Create(tutorial); err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "Inserted a single document Success")
}

func (h *TutorialHandler) ReadTutorials(c *gin.Context) {
	title := c.Query("title")
	tutorials, err := h.DB.FindAll(title)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, tutorials)
}

func (h *TutorialHandler) ReadTutorial(c *gin.Context) {
	id := c.Param("id")

	tutorials, err := h.DB.FindOne(id)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, tutorials)
}

//UpdateTutorial non test!!
func (h *TutorialHandler) UpdateTutorial(c *gin.Context) {

	tutorial := model.Tutorial{}
	if err := c.ShouldBindJSON(&tutorial); err != nil {
		c.String(500, err.Error())
		return
	}

	if err := h.DB.Update(tutorial); err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "Updated a single document Success")
}

func (h *TutorialHandler) DeleteTutorial(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.Delete(id); err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "Deleted id:"+id)
}

func (h *TutorialHandler) DeleteTutorials(c *gin.Context) {

	if err := h.DB.DeleteAll(); err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "All deleted")
}
