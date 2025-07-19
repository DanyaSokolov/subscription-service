package handler

import (
	"net/http"

	"github.com/DanyaSokolov/subscription-service/internal/model"
	"github.com/DanyaSokolov/subscription-service/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type SubscriptionHandler struct {
	Repo   *repository.SubscriptionRepository
	Logger *zap.Logger
}

func NewSubscriptionHandler(repo *repository.SubscriptionRepository, logger *zap.Logger) *SubscriptionHandler {
	return &SubscriptionHandler{Repo: repo, Logger: logger}
}

// Create godoc
// @Summary Создать подписку
// @Description Создаёт новую подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body model.SubscriptionCreateRequest true "Данные подписки"
// @Success 200 {object} model.Subscription
// @Failure 500 {object} map[string]interface{}
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		h.Logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Repo.Create(c.Request.Context(), &sub); err != nil {
		h.Logger.Error("Create failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

// GetByID godoc
// @Summary Получить подписку по ID
// @Description Возвращает подписку по её ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "ID подписки"
// @Success 200 {object} model.Subscription
// @Failure 500 {object} map[string]interface{}
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Error("Invalid UUID", zap.String("id", idParam), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sub, err := h.Repo.GetByID(c.Request.Context(), id)
	if err != nil {
		h.Logger.Error("GetByID failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription"})
		return
	}

	if sub == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

// Update godoc
// @Summary Обновить подписку
// @Description Обновляет данные подписки по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "ID подписки"
// @Param subscription body model.SubscriptionUpdateRequest true "Подписка"
// @Success 200 {object} model.Subscription
// @Failure 500 {object} map[string]interface{}
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Error("Invalid UUID", zap.String("id", idParam), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		h.Logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	sub.ID = id

	if err := h.Repo.Update(c.Request.Context(), &sub); err != nil {
		h.Logger.Error("Update failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subscription"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

// Delete godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Param id path string true "ID подписки"
// @Success 204 {string} string ""
// @Failure 500 {object} map[string]interface{}
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Error("Invalid UUID", zap.String("id", idParam), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Repo.Delete(c.Request.Context(), id); err != nil {
		h.Logger.Error("Delete failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscription"})
		return
	}

	c.Status(http.StatusNoContent)
}

// List godoc
// @Summary Список подписок
// @Description Возвращает список всех подписок
// @Tags subscriptions
// @Produce json
// @Success 200 {array} model.Subscription
// @Failure 500 {object} map[string]interface{}
// @Router /subscriptions [get]
func (h *SubscriptionHandler) List(c *gin.Context) {
	subs, err := h.Repo.List(c.Request.Context())
	if err != nil {
		h.Logger.Error("List failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list subscriptions"})
		return
	}

	c.JSON(http.StatusOK, subs)
}

// TotalCost godoc
// @Summary Общая стоимость подписок
// @Description Считает общую стоимость по фильтрам
// @Tags subscriptions
// @Produce json
// @Param        subscription body  model.TotalCostRequest  true  "Период подписок"
// @Success 200 {object} model.TotalCostResponse "successful operation"
// @Failure 500 {object} map[string]interface{}
// @Router /subscriptions/total-cost [get]
func (h *SubscriptionHandler) TotalCost(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		h.Logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	total, err := h.Repo.TotalCost(c.Request.Context(), &sub)
	if err != nil {
		h.Logger.Error("TotalCost failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total cost"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_cost": total})
}
