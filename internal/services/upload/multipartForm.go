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

	newNameFile := id.String() + "_" + file.Filename

	path := fmt.Sprintf("./assets/storage/%s", newNameFile)

	err = os.WriteFile(path, buffer, 0777)

	return newNameFile, err
}
