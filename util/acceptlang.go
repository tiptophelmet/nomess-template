package util

import (
	"fmt"
	"sort"
	"strings"
)

type languageTag struct {
	tag     string
	quality float64
}

func PickAcceptLang(langTags string) string {
	tags := ParseAcceptLang(langTags)
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].quality > tags[j].quality
	})
	return tags[0].tag
}

func ParseAcceptLang(langTags string) []languageTag {
	langTagsList := strings.Split(langTags, ",")
	tags := make([]languageTag, 0, len(langTagsList))

	for _, part := range langTagsList {
		parts := strings.Split(strings.TrimSpace(part), ";")
		tag := strings.TrimSpace(parts[0])
		quality := 1.0
		if len(parts) > 1 && strings.HasPrefix(parts[1], "q=") {
			fmt.Sscanf(strings.TrimPrefix(parts[1], "q="), "%f", &quality)
		}
		tags = append(tags, languageTag{tag, quality})
	}
	return tags
}
