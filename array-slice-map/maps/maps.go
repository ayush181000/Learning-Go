package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Instagram":           "https://instagram.com",
		"Facebook":            "https://facebook.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Facebook")
	fmt.Println(websites)
}
