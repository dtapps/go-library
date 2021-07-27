package message

type MsgType_ string

const (
	TextStr       MsgType_ = "text"
	LinkStr       MsgType_ = "link"
	MarkdownStr   MsgType_ = "markdown"
	ActionCardStr MsgType_ = "actionCard"
	FeedCardStr   MsgType_ = "feedCard"
)
