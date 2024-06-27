package resource

import (
	"fmt"
	"strings"
)

func convert(sourceFile map[string]interface{}, fieldMapping map[string]string, valueMapping map[string]string, regionMapping map[string]string) {
	convertedConfig := make(map[string]interface{})

	for inMapping, outMapping := range fieldMapping {
		splitInMappings := strings.Split(inMapping, ".")
		splitOutMappings := strings.Split(outMapping, ".")
		value := findValueRecursive(splitInMappings, sourceFile)

		value = convertValue(value, valueMapping, regionMapping)

		setValueRecursive(convertedConfig, splitOutMappings, value)
	}
	fmt.Println(convertedConfig)
}

func convertValue(valueIn string, valueMapping map[string]string, regionMapping map[string]string) string {
	if valueOut, ok := valueMapping[valueIn]; ok {
		return valueOut
	}

	if valueOut, ok := regionMapping[valueIn]; ok {
		return valueOut
	}

	return valueIn
}

func setValueRecursive(convertedConfig map[string]interface{}, outMappings []string, value string) {
	mapping := outMappings[0]

	if len(outMappings) > 1 {
		var subMap map[string]interface{}

		subMap, ok := convertedConfig[mapping].(map[string]interface{})
		if !ok {
			subMap = make(map[string]interface{})
		}

		convertedConfig[mapping] = subMap
		setValueRecursive(subMap, outMappings[1:], value)

	} else {
		convertedConfig[mapping] = value
	}
}

func findValueRecursive(inMappings []string, sourceFile map[string]interface{}) string {
	for _, mapping := range inMappings {
		if len(inMappings) > 1 {
			nestedMap, ok := sourceFile[mapping].(map[string]interface{})
			if ok {
				return findValueRecursive(inMappings[1:], nestedMap)
			}
		}
		if foundValue, ok := sourceFile[mapping].(string); ok {
			return foundValue
		}
	}
	panic("no mapping found!")
}
