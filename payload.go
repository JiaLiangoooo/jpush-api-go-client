package jpushclient

import (
	"encoding/json"
)

type PayLoad struct {
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Options      *Option     `json:"options,omitempty"`
}

func NewPushPayLoad() *PayLoad {
	pl := &PayLoad{}
	o := &Option{}
	o.ApnsProduction = false
	pl.Options = o
	return pl
}

func (pl *PayLoad) SetPlatform(pf *Platform) {
	pl.Platform = pf.Os
}

func (pl *PayLoad) SetAudience(ad *Audience) {
	pl.Audience = ad.Object
}

func (pl *PayLoad) SetOptions(o *Option) {
	pl.Options = o
}

func (pl *PayLoad) SetMessage(m *Message) {
	pl.Message = m
}

func (pl *PayLoad) SetNotice(notice *Notice) {
	pl.Notification = notice
}

func (pl *PayLoad) ToBytes() ([]byte, error) {
	content, err := json.Marshal(pl)
	if err != nil {
		return nil, err
	}
	return content, nil
}
