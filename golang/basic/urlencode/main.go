package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlStr := "https://www.google.com.hk/?redirect_url=https://www.baidu.com/"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Printf("escapeUrl: %s\n", escapeUrl)

	rawUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Printf("rawUrl: %s\n", rawUrl)
}
