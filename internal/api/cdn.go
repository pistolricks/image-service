package api

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/images"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func DeleteImage(id string) (*interface{}, error) {

	client := cloudflare.NewClient(
		option.WithAPIKey("0baf9ebe3f4276c6ca0c878d2611818fbd8ac"), // defaults to os.LookupEnv("CLOUDFLARE_API_KEY")
		option.WithAPIEmail("erikwsmith1982@gmail.com"),            // defaults to os.LookupEnv("CLOUDFLARE_EMAIL")
	)
	v1, err := client.Images.V1.Delete(
		context.TODO(),
		id,
		images.V1DeleteParams{
			AccountID: cloudflare.F("1d255d4f4a8aa1ef689a5286f80516cc"),
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", v1)

	return v1, err

}
