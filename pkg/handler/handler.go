package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sergey-timtsunyk/todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("", h.getAllList)
			lists.POST("", h.createList)
			lists.GET("/:list_id", h.getListById)
			lists.PUT("/:list_id", h.updateList)
			lists.DELETE("/:list_id", h.deleteList)

			items := lists.Group(":list_id/items")
			{
				items.GET("", h.getAllItem)
				items.POST("", h.createItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.PUT("/:item_id/done", h.doneItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
