package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wt3022/github-release-notifier/internal/db"
	"gorm.io/gorm"
)

func ListProjects(c *gin.Context, dbClient *gorm.DB) {
	/* プロジェクト一覧を取得します */
	var projects []db.Project

	if err := dbClient.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, projects)
}

func DetailProject(c *gin.Context, dbClient *gorm.DB) {
	/* プロジェクトの詳細を取得します */
	var project db.Project

	id := c.Param("id")
	if err := dbClient.First(&project, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func CreateProjects(c *gin.Context, dbClient *gorm.DB) {
	/* プロジェクトを作成します */
	var project db.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := dbClient.Create(&project)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}

	var createdProject db.Project
	if err := dbClient.First(&createdProject, project.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProject)
}

func UpdateProject(c *gin.Context, dbClient *gorm.DB) {
	/* プロジェクトを更新します */
	var project db.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dbClient.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context, dbClient *gorm.DB) {
	/* プロジェクトを削除します */
	var project db.Project

	id := c.Param("id")
	if err := dbClient.Delete(&project, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
