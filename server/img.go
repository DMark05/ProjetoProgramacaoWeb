package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getImage(imagePath string) (*os.File, error) {
	fileParts := strings.Split(imagePath, ".")
	if len(fileParts) < 2 {
		return nil, fmt.Errorf("image name didn't have an extension")
	}
	if _, err := os.Stat(imagePath); errors.Is(err, os.ErrNotExist) {

	}
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getEventImage(w http.ResponseWriter, r *http.Request) {
	imagePath := "local/images/"
	image := r.PathValue("eventId")
	if image == "" {
		http.Error(w, "must send eventId in path: %", http.StatusBadRequest)
		return
	}
	file, err := getImage(imagePath + image)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't open the image file: %s", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/"+strings.Split(imagePath, ".")[len(strings.Split(imagePath, "."))-1])
	w.Header().Set("Content-Disposition", "inline")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write image into response: %s", err), http.StatusInternalServerError)
	}
}
