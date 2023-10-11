package targetlist

// BuildHtmlTargetItem item a seleccionar
func (t TargetList) BuildTag() string {

	out := `<li data-id="` + t.FieldID + `" class="target-li-off target-li ">` //abrimos listado

	//agregamos icono
	if t.LeftMiddleText != "" {
		out += `<i name="icon" class="left-description" data-up="` + t.LeftUpText + `" data-down="` + t.LeftDownText + `" >` + t.LeftMiddleText + `</i>`
	} else if t.LeftIcon != "" {
		out += t.LeftIcon
	}

	out += `<a href="#" name="title">` + t.FieldText + `</a>`

	if t.FooterDescription != "" {
		out += `<span class="description-target">` + t.FooterDescription + `</span>`
	}

	out += `</li>`

	return out
}
