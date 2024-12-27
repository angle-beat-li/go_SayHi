package scheduler

import (
	"go_SayHi/pkg/sitemap"
	"log/slog"

	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()

	addCronFunc(c, "0 4 ? * *", func() {
		sitemap.Generate()
	})

}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	if _, err := c.AddFunc(sepc, cmd); err != nil {
		slog.Error("add cron func error", slog.Any("err", err))
	}
}
