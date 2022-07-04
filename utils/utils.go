package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"os"
	"strings"
)

// CreateFolder folder is created at the root of the project.
func CreateFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

// ImageProcessing does the mime type of the image is changed, it is compressed and then saved in the specified folder.
func ImageProcessing(buffer []byte, quality int, dirname string) (string, error) {
	filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".png"

	//converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	//if err != nil {
	//	return filename, err
	//}

	processed, err := bimg.NewImage(buffer).Process(bimg.Options{Quality: quality})
	if err != nil {
		return filename, err
	}

	writeError := bimg.Write(fmt.Sprintf("./"+dirname+"/%s", filename), processed)
	if writeError != nil {
		return filename, writeError
	}

	return filename, nil
}
