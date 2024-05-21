package file

import (
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"mime/multipart"
)

func ParseFiles(file *multipart.FileHeader) map[string]unstructured.Unstructured {

	src, _ := file.Open()

	defer src.Close()

	content, _ := io.ReadAll(src)

	obj := unstructured.Unstructured{}
	yaml.Unmarshal(content, &obj)

	abc := 1
	fmt.Println(abc)
}
