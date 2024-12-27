package sitemap

import (
	"go_SayHi/pkg/config"
	"log/slog"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
)

const (
	changefreqAlways  = "always"
	changefreqHourly  = "hourly"
	changefreqDaily   = "daily"
	changefreqWeekly  = "weekly"
	changefreqMonthly = "yearly"
	changefreqNever   = "never"
)

var building = false

// Generate
func Generate() {
	if !config.Instance.IsProd() {
		return
	}
	if building {
		slog.Info("Sitemap in building...")
		return
	}
	building = true
	defer func() {
		building = false
	}()

	sm := stm.NewSitemap(0)
	sm.SetDefaultHost(config.Instance.BaseUrl) //网站域名
	// if uploader.IsEnabledOss() {
	// 	sm.SetSitemapsHost(config.Instance.Uploader.AliyunOss.Host) // 上传至阿里云
	// } else {
	// 	sm.SetPublicPath(config.Instance.Uploader.Local.Host)
	// }

}
