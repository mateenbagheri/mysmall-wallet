package scripts

import (
	"github.com/jasonlvhit/gocron"
)

func DailySumScheduler() {
	s := gocron.NewScheduler()
	s.Every(1).Day().At("23:59").Do(DailySumCalculator)
	<-s.Start()
}
