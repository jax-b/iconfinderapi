package iconfinderapi

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
)

// DownloadIcon Downloads a specified icon into a image.Imgae
func (icofdr *Iconfinder) DownloadIcon(srcIMG Image) image.Image {
	req, err := http.NewRequest("GET", srcIMG.DownloadURL, nil)
	req.Header.Add("Authorization", "Bearer "+icofdr.apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	var img image.Image

	switch srcIMG.Format {
	case "PNG":
	case "png":
		img, _ = png.Decode(resp.Body)
		break
	case "jpg":
	case "JPG":
		img, _ = jpeg.Decode(resp.Body)
		break
	}
	resp.Body.Close()
	return img
}
