// changesClass() traces the status value in the table and sets appropriate class for it.
function changeClass(){
   let elems = document.getElementsByClassName("status");

    for (let i = elems.length - 1; i > -1; i--) {
        if ( elems[i].innerHTML === "Stopped") {
            elems[i].className = "status danger";
        } else {
            elems[i].className = "status success";
        }
    }
}

// changeState() sends POST request to the backend
//with name and status of a machine that was selected.
function changeState() {
    let table = document.getElementById("machinesTable");
    let curObj;

    if (table != null) {
        for (let i = 0; i < table.rows.length; i++) {
            table.rows[i].onclick = function() {
                curObj = tableText(this);

                let xhr = new XMLHttpRequest();
                xhr.open("POST", "/machines", false);
                xhr.setRequestHeader("Content-Type", "application/json");
                xhr.send(JSON.stringify(curObj));
            };
        }
    }
}

// tableText() extracts text from the table row.
function tableText(tableRow) {
    let name = tableRow.childNodes[1].innerText;
    let curStatus = tableRow.childNodes[7].innerText;

    return {'name': name, 'currentStatus': curStatus};
}