package jpushclient

import (
	"errors"
)

const (
	IOS      = "ios"
	ANDROID  = "android"
	WINPHONE = "winphone"
)

type Platform struct {
	Os     interface{}
	osArry []string
}

func (pf *Platform) All() {
	pf.Os = "all"
}

func (pf *Platform) Add(os string) error {
	if pf.Os == nil {
		pf.osArry = make([]string, 0, 3)
	} else {
		switch pf.Os.(type) {
		case string:
			return errors.New("platform is all")
		default:
		}
	}

	//判断是否重复
	for _, value := range pf.osArry {
		if os == value {
			return nil
		}
	}

	switch os {
	case IOS:
		fallthrough
	case ANDROID:
		fallthrough
	case WINPHONE:
		pf.osArry = append(pf.osArry, os)
		pf.Os = pf.osArry
	default:
		return errors.New("unknown platform")
	}

	return nil
}

func (pf *Platform) AddIOS() {
	pf.Add(IOS)
}

func (pf *Platform) AddAndroid() {
	pf.Add(ANDROID)
}

func (pf *Platform) AddWinphone() {
	pf.Add(WINPHONE)
}
