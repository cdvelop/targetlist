package targetlist

import "github.com/cdvelop/model"

type TargetList struct {
	Object *model.Object // ej: patient

	FieldID   string // ej: "id_client"
	FieldText string // ej: "client_name"

	LeftUpText     string
	LeftMiddleText string
	LeftDownText   string

	LeftIcon string

	FooterDescription string
}

type typeDate struct {
}
