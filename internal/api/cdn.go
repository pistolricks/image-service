package api

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/images"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"os"
)

func DeleteImage(id string) (*interface{}, error) {
	apiKey := os.Getenv("CLOUDFLARE_IMAGES_KEY")
	apiEmail := os.Getenv("VITE_IMAGE_EMAIL")
	apiAccount := os.Getenv("VITE_CLOUDFLARE_IMAGES_ACCOUNT")

	client := cloudflare.NewClient(
		option.WithAPIKey(apiKey),     // defaults to os.LookupEnv("CLOUDFLARE_API_KEY")
		option.WithAPIEmail(apiEmail), // defaults to os.LookupEnv("CLOUDFLARE_EMAIL")
	)
	v1, err := client.Images.V1.Delete(
		context.TODO(),
		id,
		images.V1DeleteParams{
			AccountID: cloudflare.F(apiAccount),
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", v1)

	return v1, err

}
