package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-multierror"
)

// larkOpt 构建飞书推送对象可选参数
type larkOpt func(l *lark)

// newLark 构建飞书推送对象
func newLark(cfg *larkConfig, opts ...larkOpt) *lark {
	l := &lark{
		cfg:    cfg,
		footer: newLarkDefaultFooter(),
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// lark 飞书推送对象
type lark struct {
	// cfg 配置
	cfg *larkConfig
	// builders 消息卡片内容构建者
	builders []larkContentBuilder
	// footer 消息卡片底部提示
	footer []*larkElement
}

// send 发送消息
func (l *lark) send(header *larkHeader, elements ...*larkElement) error {
	if l == nil || l.cfg == nil {
		return nil
	}
	cardContent := &larkMsg{
		Config: struct {
			WideScreenMode bool `json:"wide_screen_mode"`
		}{
			WideScreenMode: true,
		},
		Header:   header,
		Elements: append(l.buildAll(), elements...),
	}

	if len(l.footer) > 0 {
		cardContent.Elements = append(cardContent.Elements, l.footer...)
	}

	msgCard := &larkMsgCard{MsgType: "interactive", Card: cardContent}

	client := resty.New()
	var errs *multierror.Error
	for _, robot := range l.cfg.RobotAddr {
		_, err := client.R().SetBody(msgCard).Post(robot)
		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	return errs.ErrorOrNil()
}

// larkContentBuilder 消息卡片内容构建者
type larkContentBuilder interface {
	// build 构建新的消息卡片内容
	build(l *lark) []*larkElement
}

// withContentBuilders 添加可选内容构建者参数
func withContentBuilders(builders ...larkContentBuilder) larkOpt {
	return func(l *lark) {
		l.builders = append(l.builders, builders...)
	}
}

// buildAll 构建所有内容
func (l *lark) buildAll() []*larkElement {
	var elements = make([]*larkElement, 0, len(l.builders))
	for _, builder := range l.builders {
		newElements := builder.build(l)
		for _, ele := range newElements {
			if ele != nil {
				elements = append(elements, ele)
			}
		}
	}
	return elements
}

// newLarkDefaultFooter 构建新的默认底部
func newLarkDefaultFooter() []*larkElement {
	return []*larkElement{
		newLarkHr(),
		newLarkNote("人家才不是想要引起你的注意！绝对不是的！"),
	}
}
