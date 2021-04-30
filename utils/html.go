package utils

import "regexp"

func GetFirstImg(html string) string {
	reg, _ := regexp.Compile(`<img.+src=[\'"]([^\'"]+)[\'"].*>`)
	images := reg.FindStringSubmatch(html)

	if len(images) > 0 {
		return images[1]
	}
	return ""
}

func SubStr(html string, length int) string {
	reg, _ := regexp.Compile(`.*?<body.*?>(.*?)<\\/body>`)
	res1 := reg.ReplaceAllString(html, "$1")
	reg2, _ := regexp.Compile(`</?[a-zA-Z]+[^><]*>`)
	res2 := reg2.ReplaceAllString(res1, "")
	reg3, _ := regexp.Compile(`\n`)
	res3 := reg3.ReplaceAllString(res2, "")
	return string([]rune(res3)[0:length])
}
