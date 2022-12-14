package scripts

import (
	"database/sql"
	"log"
	"time"

	"github.com/mateenbagheri/mysmall-wallet/service"
)

func DailySumCalculator() {
	var result sql.NullFloat64
	result, err := service.SelectDailyAmount()
	log.Println("Task is being performed at:", time.Now())
	if err != nil {
		log.Fatal("something wrong with database calculations. Error:", string(err.Error()))
	}
	log.Println("today, you have gained/lost", result.Float64, "in transactions")
}
