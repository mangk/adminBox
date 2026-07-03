package httpx

import (
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"reflect"
	"strings"
	"unicode/utf8"
)

func validateHTTPCode(code int) error {
	if code < 100 || code > 599 {
		return fmt.Errorf("%w: %d", ErrInvalidHTTPCode, code)
	}
	return nil
}

func validateContentType(contentType string) error {
	if contentType == "" {
		return ErrInvalidContentType
	}

	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return fmt.Errorf("invalid content type format: %w", err)
	}

	if strings.ContainsAny(mediaType, "<>{}[]") {
		return fmt.Errorf("%w: contains dangerous characters", ErrInvalidContentType)
	}

	if charset, ok := params["charset"]; ok {
		if !isValidCharset(charset) {
			return fmt.Errorf("unsupported charset: %s", charset)
		}
	}

	return nil
}

func isValidCharset(charset string) bool {
	validCharsets := map[string]bool{
		"utf-8": true, "UTF-8": true,
		"iso-8859-1": true, "ISO-8859-1": true,
	}
	return validCharsets[charset]
}

func validateFilePath(filePath string) error {
	if filePath == "" {
		return ErrEmptyFilePath
	}

	cleanPath := filepath.Clean(filePath)
	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("%w: %s", ErrPathTraversal, filePath)
	}

	return nil
}

func validateDataSize(data interface{}) error {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)
	size := getDataSize(v)

	if size > MaxDataSize {
		return fmt.Errorf("%w: %d bytes (max: %d)", ErrDataTooLarge, size, MaxDataSize)
	}

	return nil
}

func getDataSize(v reflect.Value) int64 {
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		if v.IsNil() {
			return 0
		}
		return int64(v.Len()) * 64
	case reflect.String:
		return int64(len(v.String()))
	case reflect.Array:
		return int64(v.Len()) * 64
	default:
		return 64
	}
}

func validateHTMLTemplate(templateName string, data interface{}) error {
	if templateName == "" {
		return ErrInvalidTemplateName
	}

	if strings.Contains(templateName, "..") {
		return fmt.Errorf("%w: %s", ErrPathTraversal, templateName)
	}

	if data != nil {
		if err := checkHTMLInjection(data); err != nil {
			return err
		}
	}

	return nil
}

func checkHTMLInjection(data interface{}) error {
	v := reflect.ValueOf(data)
	return checkHTMLInjectionRecursive(v, make(map[uintptr]bool))
}

func checkHTMLInjectionRecursive(v reflect.Value, visited map[uintptr]bool) error {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil
		}
		if v.Kind() == reflect.Ptr {
			ptr := v.Pointer()
			if visited[ptr] {
				return ErrCircularReference
			}
			visited[ptr] = true
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		str := v.String()
		dangerous := []string{"<script", "javascript:", "onerror=", "onload="}
		for _, d := range dangerous {
			if strings.Contains(strings.ToLower(str), d) {
				return fmt.Errorf("%w: contains %s", ErrHTMLInjection, d)
			}
		}
		if !utf8.ValidString(str) {
			return ErrInvalidUnicode
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if err := checkHTMLInjectionRecursive(v.Index(i), visited); err != nil {
				return err
			}
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			if err := checkHTMLInjectionRecursive(key, visited); err != nil {
				return err
			}
			if err := checkHTMLInjectionRecursive(v.MapIndex(key), visited); err != nil {
				return err
			}
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.CanInterface() {
				if err := checkHTMLInjectionRecursive(field, visited); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func validateReader(reader io.Reader) error {
	if reader == nil {
		return ErrNilReader
	}
	return nil
}