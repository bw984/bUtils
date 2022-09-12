package bUtils

import (
	"path"
	"path/filepath"
	"strings"
)

func Parents(pathToTrim string, levels int) string {
	forwardSlashes := strings.Replace(pathToTrim, "\\", "/", -1)
	cleaned := filepath.Clean(forwardSlashes)
	parts := strings.Split(cleaned, "/")
	var trimmedPath string
	for i := 0; i < len(parts)-levels-1; i++ {
		if parts[i] == "" {
			trimmedPath = "/"
		} else {
			trimmedPath = filepath.Join(trimmedPath, parts[i])
		}
	}
	return trimmedPath
}

func TrimLeft(pathToTrim string, levels int) string {
	if levels < 0 {
		return ""
	}
	forwardSlashes := strings.Replace(pathToTrim, "\\", "/", -1)
	noLeadingSlash := strings.TrimLeft(forwardSlashes, "/")
	cleaned := path.Clean(noLeadingSlash)
	pathParts := strings.Split(cleaned, "/")
	var trimmedPath string
	for i := levels; i < len(pathParts); i++ {
		trimmedPath = path.Join(trimmedPath, pathParts[i])
	}
	return trimmedPath
}
