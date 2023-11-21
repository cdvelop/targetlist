package targetlist

import "github.com/cdvelop/model"

func Add(o *model.Object, ic *Item) {

	o.FrontHandler.ViewAdapter = &targetList{
		Logger:      o.Logger,
		object_name: o.ObjectName,
		Item:        ic,
	}

}

func (targetList) NameViewAdapter() string {
	return "targetList"
}
