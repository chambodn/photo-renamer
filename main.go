package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func main() {
	token := os.Getenv("DROPBOX_ACCESS_TOKEN")
	config := dropbox.Config{
		Token:    token,
		LogLevel: dropbox.LogOff, // if needed, set the desired logging level. Default is off
	}
	dbx := files.New(config)
	_, error := dbx.CreateFolderV2(files.NewCreateFolderArg("sandbox"))

	if error != nil {
		fmt.Printf("%s", error)
	}

	//fmt.Printf("%s", res.Metadata.Name)
	resu, _ := dbx.Search(files.NewSearchArg("", "*.jpg"))

	fmt.Printf("Nombre de résultats: %d\n", len(resu.Matches))
	// for _, v := range resu.Matches{
	// 	fmt.Printf("%s", v.Metadata)
	// }
	// start making API calls

	fmt.Printf("%s\n", token)
	fname := "2019-10-07 13.19.06.jpg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	fmt.Println(camModel.StringVal())

	focal, _ := x.Get(exif.FocalLength)
	numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	fmt.Printf("%v/%v", numer, denom)

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

	lat, long, _ := x.LatLong()
	fmt.Println("lat, long: ", lat, ", ", long)
}
