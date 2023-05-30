
function saveState(key, state) {
    if (state == null || state == undefined || !state) {
        localStorage.removeItem(key)
    } else {
        localStorage.setItem(key, JSON.stringify(state))
    }
}