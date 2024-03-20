package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

type UploaderMultipartForm struct{}

func (u *UploaderMultipartForm) SaveFileToDisk(file multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	buffer, _ := io.ReadAll(f)

	id, _ := uuid.NewUUID()

	path := fmt.Sprintf("./assets/storage/%s_%s", id.String(), file.Filename)

	err = os.WriteFile(path, buffer, 0777)

	return path, err
}
