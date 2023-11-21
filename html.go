package targetlist

func (t targetList) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	return `<div class="container-list-search">
	<ol class="target-list-container" data-id="` + t.ObjectName + `" onmousedown="targetListHandler(event)" ontouchstart="targetListHandler(event)">
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
// BuildHtmlTargetItem item a seleccionar

func (t targetList) BuildItemsView(all_data ...map[string]string) string {

	var out string

	for _, data := range all_data {

		out += `<li data-id="` + data[t.FieldID] + `" class="target-li-off target-li">` //abrimos listado

		out += `<div class="delete-tab"></div>`

		//agregamos icono
		if data[t.LeftMiddleText] != "" {
			out += `<i name="icon" class="left-description" data-up="` + data[t.LeftUpText] + `" data-down="` + data[t.LeftDownText] + `" >` + data[t.LeftMiddleText] + `</i>`
		} else if data[t.LeftIcon] != "" {
			out += data[t.LeftIcon]
		}

		out += `<a href="#" name="title">` + data[t.FieldText] + `</a>`

		if data[t.FooterDescription] != "" {
			out += `<span class="description-target">` + data[t.FooterDescription] + `</span>`
		}

		out += `</li>`

	}
	return out
}
