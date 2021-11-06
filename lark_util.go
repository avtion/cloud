package main

const (
	larkTagPlainText = "plain_text"
	larkTagNote      = "note"
	larkTagHr        = "hr"
	larkTagMarkdown  = "markdown"
)

// larkHeaderOpt 飞书消息片卡 Header 可选参数
type larkHeaderOpt func(header *larkHeader)

// newLarkHeader 构建新的飞书消息卡片 Header
func newLarkHeader(content string, opts ...larkHeaderOpt) *larkHeader {
	header := &larkHeader{
		Template: "turquoise",
		Title: &larkElement{
			Content: content,
			Tag:     larkTagPlainText,
		},
	}
	for _, opt := range opts {
		opt(header)
	}
	return header
}

// larkElementOpt 消息卡片内容可选参数
type larkElementOpt func(ele *larkElement)

// applyLarkElementOpts 应用消息卡片内容可选参数
func applyLarkElementOpts(ele *larkElement, opts ...larkElementOpt) *larkElement {
	for _, opt := range opts {
		opt(ele)
	}
	return ele
}

// newLarkNote 构建新的飞书消息卡片注释
func newLarkNote(content string, opts ...larkElementOpt) *larkElement {
	ele := &larkElement{
		Tag:      larkTagNote,
		Elements: []*larkElement{{Content: content, Tag: larkTagPlainText}},
	}
	return applyLarkElementOpts(ele, opts...)
}

// newLarkHr 构建新的分割线
func newLarkHr() *larkElement { return &larkElement{Tag: larkTagHr} }

// newLarkMarkdown 构建新的 Markdown 内容
func newLarkMarkdown(content string, opts ...larkElementOpt) *larkElement {
	ele := &larkElement{
		Tag:     larkTagMarkdown,
		Content: content,
	}
	return applyLarkElementOpts(ele, opts...)
}
