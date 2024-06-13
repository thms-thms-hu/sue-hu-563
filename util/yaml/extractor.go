package yaml

import (
	"fmt"
	"io"
	"mime/multipart"
	"sigs.k8s.io/yaml"
)

func ExtractContent(file *multipart.FileHeader) string {
	openedFile, _ := file.Open()
	defer openedFile.Close()

	content, _ := io.ReadAll(openedFile)

	var result map[string]interface{}
	yaml.Unmarshal(content, &result)

	foundKey := findSpecificKeyRecursive("testkey", result)
	fmt.Println(foundKey)

	return string(content)
}

func findSpecificKeyRecursive(keyword string, data map[string]interface{}) string {
	var foundKey string
	for key, value := range data {
		if subMap, ok := value.(map[string]interface{}); ok {
			foundKey = findSpecificKeyRecursive(keyword, subMap)
			if foundKey != "" {
				return foundKey
			}
		} else if key == keyword {
			return key
		}
	}
	return foundKey
}
