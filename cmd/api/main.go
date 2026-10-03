package main

import (
	"log"

	"github.com/Aufa18/ticket-booking-api/internal/config"
	"github.com/Aufa18/ticket-booking-api/internal/handler"
	"github.com/Aufa18/ticket-booking-api/internal/model"
	"github.com/Aufa18/ticket-booking-api/internal/worker"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDependencies()
	config.DB.AutoMigrate(&model.Booking{})

	go worker.StartBookingWorker()

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/admin/set-quota", handler.SetTicketQuota)
		api.POST("/book", handler.BookTicket)
	}

	log.Println("🚀 Server Ticket Booking berjalan di http://localhost:8081")
	r.Run(":8081")
}
