package targetlist

import "github.com/cdvelop/model"

func (t TargetList) ObjectVIEW() *model.Object {
	return t.Object
}

func (TargetList) ViewComponentName() string {
	return "TargetList"
}
