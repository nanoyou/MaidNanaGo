package mirai

type Event struct {
	Time     int64    `json:"time"`
	SelfID   int64    `json:"self_id"`
	PostType PostType `json:"post_type"`
}

type PostType string

const (
	PostTypeMessage   PostType = "message"
	PostTypeRequest            = "request"
	PostTypeNotice             = "notice"
	PostTypeMetaEvent          = "meta_event"
)

type MessageEvent struct {
	Event
	SubType   SubType `json:"sub_type"`
	MessageID int32   `json:"message_id"`
	UserID    int64   `json:"user_id"`
	// TODO: message是个啥
	Message    any    `json:"message"`
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	// TODO: implements
	Sender MessageSender `json:"sender"`
}

// https://docs.go-cqhttp.org/event/#%E6%B6%88%E6%81%AF%E4%B8%8A%E6%8A%A5

type SubType string

const (
	SubTypeGroup  SubType = "group"
	SubTypePublic         = "public"
)
