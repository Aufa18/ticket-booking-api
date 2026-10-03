package handler

import (
	"net/http"

	"github.com/Aufa18/ticket-booking-api/internal/config"
	"github.com/Aufa18/ticket-booking-api/internal/model"
	"github.com/Aufa18/ticket-booking-api/internal/worker"
	"github.com/gin-gonic/gin"
)

type BookingInput struct {
	UserID   string `json:"user_id" binding:"required"`
	TicketID string `json:"ticket_id" binding:"required"`
}

func SetTicketQuota(c *gin.Context) {
	config.RedisClient.Set(config.Ctx, "ticket_quota:TIX-001", 10, 0)
	c.JSON(http.StatusOK, gin.H{"message": "Kuota tiket TIX-001 diatur menjadi 10"})
}

func BookTicket(c *gin.Context) {
	var input BookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format input salah"})
		return
	}

	redisKey := "ticket_quota:" + input.TicketID

	sisaTiket, err := config.RedisClient.Decr(config.Ctx, redisKey).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengecek kuota tiket"})
		return
	}

	if sisaTiket < 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Maaf, tiket habis! Anda kalah war."})
		return
	}

	bookingData := model.Booking{
		UserID:   input.UserID,
		TicketID: input.TicketID,
	}
	worker.BookingQueue <- bookingData

	c.JSON(http.StatusOK, gin.H{
		"message":    "Selamat! Booking berhasil. Sedang dicetak sistem.",
		"sisa_tiket": sisaTiket,
	})
}
