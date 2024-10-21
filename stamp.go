package main

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"os"
	"path"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// LoadFontFace is a helper function to load the specified font file with
// the specified point size. Note that the returned `font.Face` objects
// are not thread safe and cannot be used in parallel across goroutines.
// You can usually just use the Context.LoadFontFace function instead of
// this package-level function.
func LoadEmbedFontFace(fontFile embed.FS, path string, points float64) (font.Face, error) {
	fontBytes, err := fontFile.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	return face, nil
}

func currentDir() string {
	dir, err := os.Getwd()
	//ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return dir
}

func init_context(width int, height int, fontFace font.Face) *gg.Context {
	context := gg.NewContext(width, height)
	context.SetRGB(1, 1, 1)
	context.Clear()
	context.SetRGB(0, 0, 0)
	context.SetFontFace(fontFace)
	return context
}

// Creates a triple line stamp PNG file, consisting of header, a centered date and a footer.
func CreateStamp(fontSize int, fontFile embed.FS, fontPath string, selectedTime time.Time, header string, footer string) error {
	// TODO: Allow customizing the border, as well as spacing.
	const maxCanvasSize = 4024
	const lineSpacing = 2.0

	padding := float64(fontSize)
	stroke := float64(fontSize / 10)
	radius := float64(fontSize)
	fontFace, err := LoadEmbedFontFace(fontFile, fontPath, float64(fontSize))
	if err != nil {
		return fmt.Errorf("could not load font from font file with font path %s: %s", fontPath, err)
	}
	time_format := fmt.Sprintf("%d. %s %02d", selectedTime.Day(), selectedTime.Month().String()[:3], selectedTime.Year())
	text := fmt.Sprintf("%s\n\t%s\n%s", header, footer, time_format)

	// Measure the string to get correct image size.
	context := init_context(maxCanvasSize, maxCanvasSize, fontFace)
	stringWidth, _ := context.MeasureString(strings.Split(text, "\n")[0])
	context.DrawStringWrapped(text, padding*2+stroke, padding*2+stroke, 0, 0, stringWidth, lineSpacing, gg.AlignCenter)
	measured_width, measured_height := context.MeasureMultilineString(text, lineSpacing)
	width, height := padding*4+measured_width+stroke*2, padding*4+measured_height+stroke*2

	// The resulting image & context.
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{int(width), int(height)}})
	resultContext := init_context(int(width), int(height), fontFace)

	// Draw rounded box.
	resultContext.SetColor(color.Black)
	resultContext.DrawRoundedRectangle(padding, padding, measured_width+padding*2+stroke*2, measured_height+padding*2+stroke*2, radius)
	resultContext.Fill()
	resultContext.DrawRoundedRectangle(padding+stroke, padding+stroke, measured_width+padding*2, measured_height+padding*2, radius*0.9)
	resultContext.SetColor(color.White)
	resultContext.Fill()

	// Draw text.
	resultContext.SetColor(color.Black)
	resultContext.DrawStringWrapped(text, padding*2+stroke, padding*2+stroke, 0, 0, stringWidth, lineSpacing, gg.AlignCenter)

	// Save.
	resultContext.DrawImage(img, 0, 0)
	resultContext.Clip()
	fileName := fmt.Sprintf("eingangsstempel_%d_%d_%02d.png", selectedTime.Day(), selectedTime.Month(), selectedTime.Year())

	filePath := path.Join(currentDir(), fileName)
	resultContext.SavePNG(filePath)

	fmt.Printf("Stempel unter '%s' gespeichert.\n", filePath)
	return nil
}
