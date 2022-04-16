package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserIdFormContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var inputItem data.Item
	if err := c.BindJSON(&inputItem); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.Create(userId, uint(listId), inputItem)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"item_id": itemId,
	})
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserIdFormContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var updateItem data.UpdateItemInput
	if err := c.BindJSON(&updateItem); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, uint(listId), uint(itemId), updateItem); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserIdFormContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.TodoItem.DeleteByIdAndUserId(userId, uint(listId), uint(itemId)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) doneItem(c *gin.Context) {
	userId, err := getUserIdFormContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.TodoItem.DoneItem(userId, uint(listId), uint(itemId)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserIdFormContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetByIdAndListId(userId, uint(listId), uint(itemId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

type TodoItemsResponse struct {
	Data []data.Item `json:"data"`
}

func (h *Handler) getAllItem(c *gin.Context) {
	userId, err := getUserIdFormContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, err := h.services.TodoItem.GetAllByListId(userId, uint(listId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, TodoItemsResponse{
		Data: items,
	})
}
