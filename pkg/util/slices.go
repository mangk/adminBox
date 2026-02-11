package util

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func SlicesDistinct[T comparable](slice []T) []T {
	// 使用 map 来记录已经出现的元素
	seen := make(map[T]bool)
	result := []T{}

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

func SliceColumn[T any](items []T, separator string, jsonPath ...string) (map[string][]T, error) {
	result := make(map[string][]T)

	for _, item := range items {
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		k := []string{}
		for _, p := range jsonPath {
			key := gjson.GetBytes(b, p)
			if !key.Exists() {
				return nil, fmt.Errorf("json key '%s' not found", jsonPath)
			}

			k = append(k, key.String())
		}

		sk := strings.Join(k, separator)
		result[sk] = append(result[sk], item)
	}

	return result, nil
}
