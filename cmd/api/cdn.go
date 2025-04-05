package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pistolricks/image-service/internal/api"
	"io"
	"log"
	"net/http"
	"os"
)

func imagesHandler(w http.ResponseWriter, req *http.Request) {

	apiKey := os.Getenv("CLOUDFLARE_IMAGES_KEY")
	apiEmail := os.Getenv("VITE_IMAGE_EMAIL")
	apiAccount := os.Getenv("VITE_CLOUDFLARE_IMAGES_ACCOUNT")

	requestURL := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/" + apiAccount + "/images/v1")

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("X-Auth-Email", apiEmail)
	req.Header.Set("X-Auth-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", body)

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	err = writeJSON(w, http.StatusOK, envelope{"images": string(body)}, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func imageDeleteHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	res, err := api.DeleteImage(id)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"result": res}, nil)
	if err != nil {
		log.Printf(err.Error())
		return
	}

}
