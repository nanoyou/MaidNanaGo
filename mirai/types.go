package mirai

// 所有上报
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

// 消息上报
type MessageEvent struct {
	Event
	SubType    SubType       `json:"sub_type"`
	MessageID  int32         `json:"message_id"`
	UserID     int64         `json:"user_id"`
	Message    Message       `json:"message"`
	RawMessage string        `json:"raw_message"`
	Font       int           `json:"font"`
	Sender     MessageSender `json:"sender"`
}

// TODO: message是个啥
type Message any

// TODO: implements
// 需要注意的是, sender 中的各字段是尽最大努力提供的,
// 也就是说, 不保证每个字段都一定存在,
// 也不保证存在的字段都是完全正确的(缓存可能过期).
type MessageSender any

// https://docs.go-cqhttp.org/event/#%E6%B6%88%E6%81%AF%E4%B8%8A%E6%8A%A5

type SubType string

const (
	SubTypeGroup  SubType = "group"
	SubTypePublic         = "public"
)

// 请求上报
type RequestEvent struct {
	Event
	// 请求类型
	RequestType PostType `json:"request_type"`
}

// 通知上报
type NoticeEvent struct {
	Event
	// 通知类型
	NoticeType PostType `json:"notice_type"`
}

// 元事件上报
type MetaEventType struct {
	Event
	// 元数据类型
	MetaEventType PostType `json:"meta_event_type"`
}

// 私聊消息
type PrivateMessage struct {
	Event

	// 事件数据
	MessageType  MessageType  `json:"message_type"`
	MessageEvent MessageEvent `json:"message_event"`
	TempSource   int          `json:"temp_source"`

	// 快速操作
	Reply      Message `json:"reply"`
	AutoEscape bool    `json:"auto_escape"`
}

type MessageType string
