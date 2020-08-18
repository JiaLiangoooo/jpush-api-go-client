package jpushclient

type Option struct {
	SendNo          int   `json:"sendno,omitempty"`
	TimeLive        int   `json:"time_to_live,omitempty"`
	ApnsProduction  bool  `json:"apns_production"`
	OverrideMsgId   int64 `json:"override_msg_id,omitempty"`
	BigPushDuration int   `json:"big_push_duration,omitempty"`
}

func (op *Option) SetSendno(no int) {
	op.SendNo = no
}

func (op *Option) SetTimelive(timelive int) {
	op.TimeLive = timelive
}

func (op *Option) SetOverrideMsgId(id int64) {
	op.OverrideMsgId = id
}

func (op *Option) SetApns(apns bool) {
	op.ApnsProduction = apns
}

func (op *Option) SetBigPushDuration(bigPushDuration int) {
	op.BigPushDuration = bigPushDuration
}
