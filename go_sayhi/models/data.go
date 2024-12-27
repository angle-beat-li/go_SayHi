package models

// ActionLink 站点导航
type ActionLink struct {
	Title string `json:"title"`
	Url string `json:"url"`
}

// ScoreConfig 积分配置
type ScoreConfig struct {
	PostTopicScore 		int `json:"postTopicScore"` 	// 发帖积分
	PostCommentScore 	int `json:"postCommentScore"` 	// 评论积分
	CheckInScore 		int `json:"checkInScore"` 		// 签到积分
}

type LoginMethod struct {
	Password 	bool 	`json:"password"`
	QQ			bool 	`json:"qq"`
	Github		bool 	`json:"github"`
	Qsc			bool 	`json:"qsc"`
}

// 模块配置
type ModulesConfig struct {
	Tweet 		bool 	`json:"tweet"`
	Topic 		bool 	`json:"topic"`
	Article 	bool 	`json:"article"`
}

// 配置返回结构体
type SysConfigResponse struct {
	SiteTitle  					string 				`json:"siteTitle"`
	SiteDescription 			string 				`json:"siteDescription"`
	SiteKeywords    			[]string 			`json:"siteKeywords"`
	SiteNavs					[]ActionLink 		`json:"siteNavs"`
	SiteNotification 			string 				`json:"siteNotification"`
	RecommendTags				[]string 			`json:"recommendTags"`
	UrlRedirect					bool 				`json:"urlRedirect"`
	ScoreConfig					ScoreConfig 		`json:"scoreConfig"`
	LoginMethod					LoginMethod			`json:"loginMethod"`
	DefaultNodeId				int64				`json:"defaultNodeId"`
	ArticlePending				bool 				`json:"articlePending"`
	TopicCaptcha				bool 				`json:"topicCaptcha"`
	UserObserveSeconds			int64				`json:"userObserveSeconds"`
	TokenExpireDays				int					`json:"tokenExpireDays"`
	CreateTopicEmailVerified	bool 				`json:"createTopicEmailVerified"`
	CreateArticleEmailVerified	bool 				`json:"createArticleEmailVerified"`
	CreateCommentEmailVerified	bool 				`json:"createCommentEmailVerified"`
	EnableHiddeContent			bool 				`json:"enableHiddeContent"`
	Modules						ModulesConfig		`json:"modules"`
	EmailWhitelist				[]string			`json:"emailWhitelist"`
}

