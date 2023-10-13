

function TargetListSelected(e) {
    console.log("EVENTO:",e);

	if (e.target && e.target.tagName === 'LI') {
		console.log("SELECCIONADO:",e.target);
		

        // userViewComponentClicked(optionsContainer.dataset.id, target.dataset.id);
	}
}