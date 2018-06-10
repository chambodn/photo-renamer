package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/chambodn/photo-renamer/log"
	"github.com/xiam/exif"
)

func main() {
	data, err := exif.Read("/Users/vidou/Desktop/104D5100/31-10-2016_DSC_0757.JPG")
	if err != nil {
		fmt.Printf("Something goes wrong")
	}
	//Date and Time = 2016:10:31 06:55:35
	//only interested in Date and time at this stage:
	dateAndTime := data.Tags["Date and Time"]
	log.Logger.Info("Extract Metadata", zap.String("DateTime", dateAndTime))

	value := strings.Split(dateAndTime, " ")
	log.Logger.Info("Split Time and Date", zap.String("Date", value[0]), zap.String("Time", value[1]))

	sDate := strings.Replace(value[0], ":", "-", -1)
	log.Logger.Info("We need to replace ':' by '-", zap.String("Date", sDate))

	if _, err := os.Stat(sDate); os.IsNotExist(err) {
		// directory sDate does not exist
		log.Logger.Warn("Directory does not exit", zap.String("Directory", sDate))
	}

	if _, err := os.Stat("/path/to/whatever"); err == nil {
		log.Logger.Warn("Directory exists", zap.String("Directory", sDate))
	}
	//layout := "2016:10:31 06:55:35"

}

func parseDate(date string) time.Time {
	log.Logger.Info("salut")

	return time.Now()
}
