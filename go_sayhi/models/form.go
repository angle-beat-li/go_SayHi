package models

import (
	"go_SayHi/models/constants"
	"go_SayHi/pkg/common"
	"log/slog"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/common/jsons"
	"github.com/mlogclub/simple/common/strs"
	"github.com/mlogclub/simple/web/params"
	"github.com/tidwall/gjson"
)

type CreateTopicForm struct {
	Type        constants.TopicType `json:"type"`
	CaptchaId   string              `json:"captchaId"`
	CaptchaCode string              `json:"captchaCode"`
	NodeId      int64               `json:"nodeId"`
	Title       string              `json:"title"`
	Content     string              `json:"content"`
	HideContent string              `json:"hideContent"`
	Tags        []string            `json:"tags"`
	ImageList   []ImageDTO          `json:"imageList"`
	UserAgent   string              `json:"userAgent"`
	Ip          string              `json:"ip"`
}

type CreateArticleForm struct {
	Title       string
	Summary     string
	Content     string
	ContentType string
	Cover       *ImageDTO
	Tags        []string
	SourceUrl   string
}

// 发表评论
type CreateCommentForm struct {
	EntityType string     `json:"entityType"`
	EntityId   int64      `json:"entityId"`
	Content    string     `json:"content"`
	ImageList  []ImageDTO `json:"imageList"`
	QuoteId    int64      `json:"quoteId"`
	UserAgent  string     `json:"userAgent"`
	Ip         string     `json:"ip"`
}

type ImageDTO struct {
	Url string `json:"url"`
}

func GetCreateTopicForm(ctx iris.Context) CreateTopicForm {
	contentType := ctx.GetHeader("content-Type")

	var form *CreateTopicForm
	if contentType == "application/json" {
		if err := ctx.ReadJSON(&form); err != nil {
			slog.Error(err.Error(), slog.Any("err", err))
		}
	} else {
		form = &CreateTopicForm{
			Type:        constants.TopicType(params.FormValueIntDefault(ctx, "type", int(constants.TopicTypeTopic))),
			CaptchaId:   params.FormValue(ctx, "captchaId"),
			CaptchaCode: params.FormValue(ctx, "captchaCode"),
			NodeId:      params.FormValueInt64Default(ctx, "nodeId", 0),
			Title:       strings.TrimSpace(params.FormValue(ctx, "title")),
			Content:     strings.TrimSpace(params.FormValue(ctx, "content")),
			HideContent: strings.TrimSpace(params.FormValue(ctx, "hidecontent")),
			Tags:        params.FormValueStringArray(ctx, "tags"),
			ImageList:   GetImageList(ctx, "imageList"),
			UserAgent:   common.GetUserAgent(ctx.Request()),
			Ip:          common.GetRequestIp(ctx.Request()),
		}
	}
	return *form
}

func GetCreateCommentForm(ctx iris.Context) CreateCommentForm {
	form := CreateCommentForm{
		EntityType: params.FormValue(ctx, "entityType"),
		EntityId:   params.FormValueInt64Default(ctx, "entityId", 0),
		Content:    strings.TrimSpace(params.FormValue(ctx, "content")),
		ImageList:  GetImageList(ctx, "imageList"),
		QuoteId:    params.FormValueInt64Default(ctx, "quoteId", 0),
		UserAgent:  common.GetUserAgent(ctx.Request()),
		Ip:         common.GetRequestIp(ctx.Request()),
	}
	return form
}

func GetCreateArticleForm(ctx iris.Context) CreateArticleForm {
	var (
		title   = ctx.PostValue("title")
		summary = ctx.PostValue("summary")
		content = ctx.PostValue("content")
		tags    = params.FormValueStringArray(ctx, "tags")
		cover   = GetImageDTO(ctx, "cover")
	)

	return CreateArticleForm{
		Title:       title,
		Summary:     summary,
		Content:     content,
		ContentType: constants.ContentTypeMarkdown,
		Cover:       cover,
		Tags:        tags,
	}
}

func GetImageList(ctx iris.Context, paramName string) []ImageDTO {
	imageListStr := params.FormValue(ctx, paramName)
	var imageList []ImageDTO
	if strs.IsNotBlank(imageListStr) {
		ret := gjson.Parse(imageListStr)
		if ret.IsArray() {
			for _, item := range ret.Array() {
				url := item.Get("url").String()
				imageList = append(imageList, ImageDTO{
					Url: url,
				})
			}
		}
	}
	return imageList
}

func GetImageDTO(ctx iris.Context, paramName string) (image *ImageDTO) {
	str := params.FormValue(ctx, paramName)
	if strs.IsBlank(str) {
		return
	}
	if err := jsons.Parse(str, &image); err != nil {
		slog.Error(err.Error(), slog.Any("err", err))
	}
	return
}
