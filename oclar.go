package oclar

import "fmt"

type Style []string

func (style Style) Render(text string) string {

	var stylecodes string = ""
	for _, code := range style {
		if _, ok := easycodes[code]; ok {
			code = easycodes[code]
		}
		stylecodes += ";" + code
	}
	text = "\u001b[" + stylecodes + "m" + text + "\u001b[0m"
	return text
}
func (style *Style) DeleteRule(rule string) {
	b := (*style)[:0]
	for _, item := range *style {
		if item != rule {
			b = append(b, item)
		}
	}
	*style = b
}
func (style *Style) AddRule(rule string) {
	*style = append(*style, rule)
}
func main() {
	s := Style{"greenfg", "underline"}
	fmt.Println(s.Render("Hello"))
	fmt.Println(s)
	s.DeleteRule("underline")
	fmt.Println(s.Render("Hello"))
	fmt.Println(s)
}
