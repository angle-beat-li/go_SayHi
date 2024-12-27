package markdown

import (
	"go_SayHi/pkg/text"
	"sync"

	"github.com/88250/lute"
	"github.com/mlogclub/simple/common/strs"
)

var (
	engine *lute.Lute
	once   sync.Once
)

func getEngine() *lute.Lute {
	once.Do(func() {
		engine = lute.New(func(lute *lute.Lute) {
			lute.SetSanitize(true)
			lute.SetGFMTaskListItem(true)
		})
	})
	return engine
}

func ToHTML(markdownStr string) string {
	if strs.IsBlank(markdownStr) {
		return ""
	}
	return getEngine().MarkdownStr("", markdownStr)
}

func GetSummary(markdownStr string, summaryLen int) string {
	htmlStr := ToHTML(markdownStr)
	return text.GetSummary(htmlStr, summaryLen)
}
