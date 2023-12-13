package targetlist

import "github.com/cdvelop/model"

type targetList struct {
	*model.Object
	*Item
	reset model.CallJsFunWithParameters
}

type Item struct {
	FieldID   string // ej: "id_client"
	FieldText string // ej: "client_name"

	LeftUpText     string
	LeftMiddleText string
	LeftDownText   string

	LeftIcon string

	FooterDescription string
}
