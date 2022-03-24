// changesClass() traces the status value in the table and sets appropriate class for it.
function changeClass(){
   let elems = document.getElementsByClassName("status")
    for (let i = elems.length - 1; i > -1; i--) {
        if ( elems[i].innerHTML === "Stopped") {
            elems[i].className = "status danger"
        } else {
            elems[i].className = "status success"
        }
    }
}