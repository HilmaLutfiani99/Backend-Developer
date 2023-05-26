package main

import (
	"fmt"
	"strings"
)

func searchKeywords(text string, keywords ...string) {
	for _, keyword := range keywords {
		count := strings.Count(text, keyword)
		for i := 0; i < count; i++ {
			fmt.Printf("%s - ", keyword)
		}
	}
}

func main() {
	text := "Berikut adalah kisah sang gajah. Sang gajah memiliki teman serigala bernama DoeSang. Gajah sering dibela oleh serigala ketika harimau mendekati gajah."
	keywords := []string{"sang gajah", "serigala", "harimau"}

	searchKeywords(text, keywords...)
}
