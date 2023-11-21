package targetlist

import "github.com/cdvelop/model"

func Add(o *model.Object, ic *Item) {

	o.FrontHandler.ViewAdapter = &targetList{
		Object: o,
		Item:   ic,
	}

}

func (targetList) NameViewAdapter() string {
	return "targetList"
}
