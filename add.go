package targetlist

import "github.com/cdvelop/model"

func Add(o *model.Object, ic *Item) {

	o.FrontHandler.ObjectViewHandler = &targetList{
		Object: o,
		Item:   ic,
		reset: model.CallJsOptions{
			NameJsFunc: "resetTargetList",
			Enable:     true,
			Params:     map[string]any{},
		},
	}

}

func (targetList) NameViewAdapter() string {
	return "targetList"
}
