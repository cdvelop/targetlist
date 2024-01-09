

function deleteListSelected(target) {
    // console.log('BORRAR ELEMENTO:', target);

    const obj_name = target.parentNode.dataset.id;
    // console.log('obj_name ', obj_name);

    // console.log('ID', target.dataset.id);

    let ids = [target.dataset.id]

    deleteObject(obj_name, ids)
}


function deleteTargetList(o) {
    // console.log("deleteTargetList", o)
    const container = document.querySelector(o.query)
    // console.log("container", container)
    if (container != undefined) {
        let item = container.querySelector('[data-id="' + o.id + '"]')
        if (item != undefined) {
            item.remove();
        }
    }
}