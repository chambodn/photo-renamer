package main

import (
	"fmt"

	"github.com/xiam/exif"
)

func main() {
	data, err := exif.Read("/Users/vidou/Desktop/104D5100/31-10-2016_DSC_0757.JPG")
	if err != nil {
		fmt.Printf("Something goes wrong")
	}
	//Date and Time = 2016:10:31 06:55:35
	//only interested in Date and time at this stage:
	fmt.Printf("Youhou = %s\n\n", data.Tags["Date and Time"])

	//layout := "2016:10:31 06:55:35"

}
,