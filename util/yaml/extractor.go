package yaml

import (
	"io"
	"mime/multipart"
)

func ExtractContent(file *multipart.FileHeader) string {
	openedFile, _ := file.Open()
	defer openedFile.Close()

	content, _ := io.ReadAll(openedFile)
	return string(content)
}
