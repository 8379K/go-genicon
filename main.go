package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/jakobvarmose/go-qidenticon"
)

const iconSize = 400

var iconSettings = qidenticon.DefaultSettings()

// generateIcon アイコン画像を生成します
func generateIcon(salt string) image.Image {
	return qidenticon.Render(qidenticon.Code(salt), iconSize, iconSettings)
}

func timestamp() string {
	t := time.Now()
	return fmt.Sprintf("%02d%02d%02d_%02d%02d%02d",
		t.Year()%100, t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
}

func main() {
	ts := timestamp()
	for i := 0; i < 100; i++ {
		fpath := fmt.Sprintf("./img/image_%s(%d).png", ts, i+1)
		renderimage := generateIcon(fpath)

		f, _ := os.Create(fpath)
		png.Encode(f, renderimage)
		f.Close()
	}
}
