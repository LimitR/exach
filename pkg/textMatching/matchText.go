package textmatching

import (
	"regexp"
)

func MatchTextToHTML(text string) string {
	exp := regexp.MustCompile(`<br>`)
	text = exp.ReplaceAllString(text, "(\r\n)|(\n)")

	exp = regexp.MustCompile(`(&gt;&gt;([\d]+))`)
	text = exp.ReplaceAllString(text, "<a href=\"#c$2\" onclick=\"comment_css('c$2');\" onmouseover=\"preview_comment(event, '$2');\">&gt;&gt;$2</a>")

	exp = regexp.MustCompile(`\[b\](.+?)\[/b\]`)
	text = exp.ReplaceAllString(text, "<b>$1</b>")

	exp = regexp.MustCompile(`\[quote\](.+?)\[/quote\]`)
	text = exp.ReplaceAllString(text, "<span class=\"quote\">$1</span>")

	exp = regexp.MustCompile(`\[b\](.+?)\[/b\]`)
	text = exp.ReplaceAllString(text, "<b>$1</b>")

	exp = regexp.MustCompile(`\[i\](.+?)\[/i\]`)
	text = exp.ReplaceAllString(text, "<i>$1</i>")

	exp = regexp.MustCompile(`\[u\](.+?)\[/u\]`)
	text = exp.ReplaceAllString(text, "<ins>$1</ins>")

	exp = regexp.MustCompile(`\[s\](.+?)\[/s\]`)
	text = exp.ReplaceAllString(text, "<del>$1</del>")

	exp = regexp.MustCompile(`\[spoiler\](.+?)\[/spoiler\]`)
	text = exp.ReplaceAllString(text, "<span class=\"spoiler\">$1</span>")

	exp = regexp.MustCompile(`\[smile\]([1-9]|[123][0-9]|[4][0-3])\[/smile\]`)
	text = exp.ReplaceAllString(text, "<img src=\"/public/img/smiles/$1.gif\" alt=\"\" />")

	return text
}

func LinkToImage(nameImg string) string {
	return `<a href="/public/storage/` + nameImg + `" class="text_image"><img style="width: 300px" onclick="show_image(this,'200px'); return false" src="/public/storage/` + nameImg + `" alt=""></a>`
}
