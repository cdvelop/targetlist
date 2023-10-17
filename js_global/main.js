

function TargetListSelected(e) {
	let target = e.target

	const tagname = target.tagName.toLowerCase()
	// console.log("tagname:", tagname)


	if (tagname == 'span' || tagname == 'i') {
		target = e.target.parentNode
	}

	if (target && tagname != 'ol') {

		userViewComponentClicked(target.parentNode.dataset.id, target.dataset.id);
		OnOffSelList(target)
	}

	// console.log("target:", target);
}

// selección actual, anterior
let targetOLD;
function OnOffSelList(t) {
	if (targetOLD == undefined) {
		t.classList.add("target-li-on");
		t.classList.remove("target-li-off");
		targetOLD = t;
	} else {
		targetOLD.classList.remove("target-li-on");
		targetOLD.classList.add("target-li-off");
		t.classList.add("target-li-on");
		t.classList.remove("target-li-off");
		targetOLD = t;
	}
};


function TargetListClicking(module, id) {
	const container = module.querySelector(`ol.target-list-container`);

	const li = container.querySelector(`li[data-id="` + id + `"]`);
	if (li == null) {
		return "error objeto id:" + id + " no encontrado para hacer click"
	}

	// console.log("ID OK:", li)
	if (li) {
		li.click();
	}

};