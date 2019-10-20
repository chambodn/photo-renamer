package main

import (
	"log"
	"os"

	"github.com/chambodn/photo-renamer/dropbox"
	"github.com/chambodn/photo-renamer/model"
)

func main() {
	token := os.Getenv("DROPBOX_ACCESS_TOKEN")
	d := dropbox.NewFiles(dropbox.NewConfig(token))

	out, err := d.Search(&dropbox.SearchInput{
		Query: "*.jpg",
		Options: &dropbox.SearchOptions{
			SearchPath:     "",
			MaxResults:     10,
			FileStatus:     "active",
			FilenameOnly:   true,
			FileCategories: []string{"image"},
		},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, entry := range out.Matches {
		log.Printf("%s", entry.Metadata.PathLower)

		//https://making.pusher.com/alternatives-to-sum-types-in-go/
		// switch e := entry.Metadata(type) {
		// 	case
		// }
	}

	// // Create a Resty Client
	// client := resty.New()

	// resp, err := client.R().
	// 	SetHeader("Accept", "application/json").
	// 	SetHeader("Content-Type", "application/json").
	// 	SetAuthToken(token).
	// 	SetBody(NewGetMetadataArg("/Photos")).
	// 	EnableTrace().

	// 	Post("https://api.dropboxapi.com/2/files/get_metadata")

	// // Explore response object
	// fmt.Println("Response Info:")
	// fmt.Println("Error      :", err)
	// fmt.Println("Status Code:", resp.StatusCode())
	// fmt.Println("Status     :", resp.Status())
	// fmt.Println("Time       :", resp.Time())
	// fmt.Println("Received At:", resp.ReceivedAt())
	// fmt.Println("Body       :\n", resp)
	// fmt.Println()

}

// func getMetadata(request GetMetadataArg) (*GetMetadataResponse, error) {

// 	return NewGetMetadataArg(""), nil
// }

// NewGetMetadataArg returns a new GetMetadataArg instance
func NewGetMetadataArg(path string) *model.GetMetadataArg {
	s := new(model.GetMetadataArg)
	s.Path = path
	s.IncludeMediaInfo = false
	s.IncludeDeleted = false
	s.IncludeHasExplicitSharedMembers = false
	return s
}
