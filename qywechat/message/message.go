package message

type Message struct {
	MsgType  MsgType_  `json:"msgtype"`
	Text     Text_     `json:"text"`
	Markdown Markdown_ `json:"markdown"`
	News     News_     `json:"link"`
	File     File_     `json:"file"`
}

type MsgType_ string

const (
	TextStr     MsgType_ = "text"
	NewsStr     MsgType_ = "news"
	MarkdownStr MsgType_ = "markdown"
	fileStr     MsgType_ = "file"
)

// Text_ text类型
type Text_ struct {
	Content             string   `json:"content"`               // 文本内容，最长不超过2048个字节，必须是utf8编码
	MentionedList       []string `json:"mentioned_list"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
	MentionedMobileList []string `json:"mentioned_mobile_list"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
}

// Markdown_ markdown类型
type Markdown_ struct {
	Content string `json:"content"` // markdown内容，最长不超过4096个字节，必须是utf8编码
}

// News_ news类型
type News_ struct {
	Articles []articles `json:"articles"` // 图文消息，一个图文消息支持1到8条图文
}

//  articles
type articles struct {
	Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断
	Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`         // 点击后跳转的链接。
	Picurl      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
}

// File_ file类型
type File_ struct {
	MediaId string `json:"media_id"` // 文件id，通过下文的文件上传接口获取
}
