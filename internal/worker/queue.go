package worker

import (
	"fmt"
	"log"

	"github.com/Aufa18/ticket-booking-api/internal/config"
	"github.com/Aufa18/ticket-booking-api/internal/model"
)

var BookingQueue = make(chan model.Booking, 5000)

func StartBookingWorker() {
	fmt.Println("👷 Worker siap memproses antrean pesanan tiket ke Database...")

	for booking := range BookingQueue {
		err := config.DB.Create(&booking).Error
		if err != nil {
			log.Printf("Gagal menyimpan pesanan %s ke DB: %v\n", booking.UserID, err)
		} else {
			log.Printf("✅ Tiket berhasil dicetak ke DB untuk User: %s\n", booking.UserID)
		}
	}
}
