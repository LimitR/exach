async function updateForClick(classNameForEvent, classNameForUpdate, fnToHtml) {
    const element = document.querySelector(classNameForEvent)
    element.addEventListener('click', async (event) => {
        const elementToUpdate = document.querySelector(classNameForUpdate)
        const result = await fnToHtml()
        if (result) {
            elementToUpdate.innerHTML = result
        }
    })
}


async function updateToHtml(classNameForUpdate, fnToHtml) {
    const elementToUpdate = document.querySelector(classNameForUpdate)
    const result = await fnToHtml()
    if (result) {
        elementToUpdate.innerHTML = result
    }
}

async function addToHtml(classNameForUpdate, fnToHtml) {
    const elementToUpdate = document.querySelector(classNameForUpdate)
    const result = await fnToHtml()
    if (result) {
        const newElement = document.createElement("div");
        newElement.innerHTML = result
        elementToUpdate.appendChild(newElement)
    }
}