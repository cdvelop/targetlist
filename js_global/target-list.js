
function targetListHandler(e) {
	e.stopPropagation();
	const tagname = e.target.tagName.toLowerCase()
	// console.log("tagname:", tagname);
	if (tagname === "li") {
		targetHandler(e.target, targetListSelected, deleteListSelected)
	}
}


function targetListSelected(target) {
	// console.log("targetListSelected:", target);
	objectClickedUI(target.parentNode.dataset.id, target.dataset.id);
	OnOffSelList(target)
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


function TargetListObjectClicking(module, id) {
	const container = module.querySelector(`ol.target-list-container`);

	const target = container.querySelector(`li[data-id="` + id + `"]`);
	if (target == null) {
		return "error objeto id:" + id + " no encontrado para hacer click"
	}
	// console.log("TargetListObjectClicking:", target)
	// if (li) {
	// 	li.click();
	// }
	targetListSelected(target)
};


function resetTargetList(p) {
	const list = document.querySelector(p.query)
	const onoff = p.enable;

	// console.log("SAVE MODE: ", lista);
	let remove = "disable";
	let add = "target-li-off";

	if (!onoff) {
		remove = "target-li-off", add = "disable";
	};
	// lista disable
	let li = list.querySelectorAll('li');
	for (let i = 0; i < li.length; ++i) {
		// console.log("SAVE MODE: ",MAIN_SAVE_FUNCTION);
		if (!onoff) {
			li[i].classList.remove("target-li-on");
		}
		li[i].classList.remove(remove);
		li[i].classList.add(add);
	}
};


function targetListObjectCount(opt) {
	let list = document.querySelector(opt.query)

	// Obtener la lista de elementos li dentro del ol
	var elements = list.getElementsByTagName('li');

	// Contar el número de elementos li
	return elements.length
}