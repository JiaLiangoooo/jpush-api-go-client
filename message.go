package jpushclient

type Message struct {
	Content     string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func (msg *Message) SetContent(c string) {
	msg.Content = c

}

func (msg *Message) SetTitle(title string) {
	msg.Title = title
}

func (msg *Message) SetContentType(t string) {
	msg.ContentType = t
}

func (msg *Message) AddExtras(key string, value interface{}) {
	if msg.Extras == nil {
		msg.Extras = make(map[string]interface{})
	}
	msg.Extras[key] = value
}
