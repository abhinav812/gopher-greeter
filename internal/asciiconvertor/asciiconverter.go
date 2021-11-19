package asciiconvertor

import (
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg" // import to process jpeg images
	_ "image/png"  // import to process png messages
)

//ConvertToAsii - converts image file to an ascii-art and returns the string representing it.
func ConvertToAsii(imageFileName string) string {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 70
	convertOptions.FixedHeight = 40
	convertOptions.Reversed = true
	convertOptions.Colored = false

	// Create the image converter
	converter := convert.NewImageConverter()
	asciiText := converter.ImageFile2ASCIIString(imageFileName, &convertOptions)

	return asciiText
}
