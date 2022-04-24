package utils

import (
	"fmt"
	"image"
	"io/ioutil"
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type User struct {
	Name string
	Ping int
}

func getPing(ping int, dir string) string {
	if ping < 0 {
		return filepath.Join(dir, "/assets/signal_0.png")
	}
	if ping <= 150 {
		return filepath.Join(dir, "/assets/signal_5.png")
	}
	if ping <= 300 {
		return filepath.Join(dir, "/assets/signal_4.png")
	}
	if ping <= 600 {
		return filepath.Join(dir, "/assets/signal_3.png")
	}
	if ping <= 1000 {
		return filepath.Join(dir, "/assets/signal_2.png")
	}
	return filepath.Join(dir, "/assets/signal_1.png")
}

func getAvatarImg(user string) (image.Image, error) {
	url := "https://mc-heads.net/avatar/" + user + "/32"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	m, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return m, nil

}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

var ()

func DrawTablist(usersAndPing []User) *image.RGBA {
	dir, err := GetCurrentDirectory()
	if err != nil {
		fmt.Println(err)
	}
	width := 600
	height := 350
	// img := image.NewRGBA(image.Rect(0, 0, width, 350))

	backgroundImg, err := getImageFromFilePath(filepath.Join(dir, "/assets/tabgb.png"))

	if err != nil {
		fmt.Println(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), backgroundImg, image.ZP, draw.Over)

	fontBytes, err := ioutil.ReadFile(filepath.Join(dir, "/assets/mc.otf"))
	if err != nil {
		fmt.Println(err)
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		fmt.Println(err)
	}

	fontFace, _ := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: fontFace,
	}

	drawBlock := func(x, y int, user string, ping int) {
		avatarImg, err := getAvatarImg(user)
		if err != nil {
			fmt.Println(err)
		}

		pingImg, err := getImageFromFilePath(getPing(ping, dir))
		if err != nil {
			fmt.Println(err)
		}

		draw.Draw(img, image.Rect(x+10, y+10, x+41, y+300), avatarImg, image.Point{}, draw.Src)
		draw.Draw(img, image.Rect(x+10, y+12, x+41, y+300), pingImg, image.Point{}, draw.Src)

		d.Dot = fixed.Point26_6{
			X: fixed.I(x + 45),
			Y: fixed.I(y + 36),
		}

		//X left-right
		//Y up-down

		d.DrawString(user)

	}

	fillTab := func(users []User) {

		y := 0
		x := 0

		for _, user := range users {

			if y > 300 {
				x = x + 138
				y = 0
			}

			if len(user.Name) > 10 {
				user.Name = user.Name[:10] + "..."
			}

			drawBlock(x, y, user.Name, user.Ping)

			y = y + 33
		}

	}

	fillTab(usersAndPing)
	return img
}

func GenerateTablist(users []User, w http.ResponseWriter) {

	img := DrawTablist(users)

	w.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(w, img, nil)

}
