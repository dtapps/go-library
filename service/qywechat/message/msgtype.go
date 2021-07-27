package message

type MsgType_ string

const (
	TextStr     MsgType_ = "text"
	NewsStr     MsgType_ = "news"
	MarkdownStr MsgType_ = "markdown"
	fileStr     MsgType_ = "file"
)
