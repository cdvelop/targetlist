package targetlist

// BuildHtmlTargetItem item a seleccionar
func (t TargetList) BuildItemView(all_data ...map[string]string) string {

	out := `<li data-id="` + t.FieldID + `" class="target-li-off target-li">` //abrimos listado

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

func (t TargetList) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	return `<div class="container-list-search">
	<ol class="target-list-container" data-id="` + t.Object.Name + `" onclick="TargetListSelected(event)">
	</ol>
	</div>
	<div id="device-search-form" class="search-container">
	<label for="filter-object">
	<svg aria-hidden="true" focusable="false" class="form-btn">
	<use xlink:href="#icon-search" />
	</svg>
	</label>
	<input type="search" id="filter-object" title="letras números - permitidos max 50 caracteres"
	pattern="^[A-Za-zÑñ 0-9-]{1,20}$">
	</div>`
}

// <li data-id="1672322764831794600" class="target-li-off target-li"><a href="#" name="title">eco abdominal</a><span class="description-target">dra. sonia sandoval</span></li>
// <li data-id="1672322764831794600" class="target-li-off target-li"><a href="#" name="title">eco abdominal</a><span class="description-target">dra. sonia sandoval</span></li>
