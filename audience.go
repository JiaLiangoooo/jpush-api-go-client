package jpushclient

const (
	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"
)

type Audience struct {
	Object   interface{}
	audience map[string][]string
}

func (aud *Audience) All() {
	aud.Object = "all"
}

func (aud *Audience) SetID(ids []string) {
	aud.set(ID, ids)
}

func (aud *Audience) SetTag(tags []string) {
	aud.set(TAG, tags)
}

func (aud *Audience) SetTagAnd(tags []string) {
	aud.set(TAG_AND, tags)
}

func (aud *Audience) SetAlias(alias []string) {
	aud.set(ALIAS, alias)
}

func (aud *Audience) set(key string, v []string) {
	if aud.audience == nil {
		aud.audience = make(map[string][]string)
		aud.Object = aud.audience
	}

	//v, ok = aud.audience[key]
	//if ok {
	//	return
	//}
	aud.audience[key] = v
}
