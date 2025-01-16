package FileDetector

import (
	"strings"
)

func Scan_File(file string) string {
	xfile := strings.Index(file, ".")
	if xfile == -1 {
		return ""
	}
	filename := strings.Split(file, ".")
	type_file := filename[1]
	return type_file
}
func Type_File(typefile string) string {
	if strings.ToLower(typefile) == "go" {
		return "Go📦"
	} else if strings.ToLower(typefile) == "sum" {
		return "Sum📦"
	} else if strings.ToLower(typefile) == "mod" {
		return "Mod📦"
	} else if strings.ToLower(typefile) == "mp3" || strings.ToLower(typefile) == "wav" {
		return "🎵"
	} else if strings.ToLower(typefile) == "mp4" || strings.ToLower(typefile) == "mkv" || strings.ToLower(typefile) == "gif" || strings.ToLower(typefile) == "jpg" || strings.ToLower(typefile) == "png" {
		return "🖼️"
	} else {
		return "📄"
	}
}
