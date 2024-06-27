package yaml_configs

import (
	"io"
	"mime/multipart"
)

func ExtractContent(file *multipart.FileHeader) []byte {
	openedFile, _ := file.Open()
	defer openedFile.Close()

	content, _ := io.ReadAll(openedFile)

	return content
}

func FindValueByKey(keyword string, data map[string]interface{}) string {
	for key, value := range data {
		if subMap, ok := value.(map[string]interface{}); ok {
			if result := FindValueByKey(keyword, subMap); result != "" {
				return result
			}
		} else if key == keyword {
			if strValue, ok := value.(string); ok {
				return strValue
			}
		}
	}
	return ""
}
