package main

// larkConfig 飞书推送配置
type larkConfig struct {
	// RobotAddr 机器人地址
	RobotAddr []string
}

type (
	// larkMsgCard 飞书推送卡片
	larkMsgCard struct {
		MsgType string   `json:"msg_type"`
		Card    *larkMsg `json:"card"`
	}
	// larkMsg 飞书推送卡片内容
	larkMsg struct {
		Config struct {
			WideScreenMode bool `json:"wide_screen_mode"`
		} `json:"config"`
		Header   *larkHeader    `json:"header"`
		Elements []*larkElement `json:"elements"`
	}

	// larkHeader 飞书推送标题
	larkHeader struct {
		Template string       `json:"template"`
		Title    *larkElement `json:"title"`
	}

	// larkElement 飞书推送元素
	larkElement struct {
		Content string `json:"content,omitempty"`
		Href    *struct {
			UrlVal *struct {
				Url string `json:"url"`
			} `json:"urlVal"`
		} `json:"href,omitempty"`
		Tag   string `json:"tag"`
		Extra *struct {
			Alt    *larkElement `json:"alt"`
			ImgKey string       `json:"img_key"`
			Tag    string       `json:"tag"`
		} `json:"extra,omitempty"`
		Text     *larkElement   `json:"text,omitempty"`
		Elements []*larkElement `json:"elements,omitempty"`
	}
)
