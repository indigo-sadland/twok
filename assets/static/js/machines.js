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
    let curObj;
    let table = document.getElementById("machinesTable");
    let spinner = document.getElementById("spinner");

    if (table != null) {
        spinner.style.visibility = "visible";
        for (let i = 0; i < table.rows.length; i++) {
            table.rows[i].onclick = function() {
                curObj = tableText(this);

                let http = new XMLHttpRequest();
                http.open("POST", "/machines", false);
                http.setRequestHeader("Content-Type", "application/json");
                http.send(JSON.stringify(curObj));

                // Wait for response from the server to hide the spinner.
                http.onreadystatechange = function () {
                    if(http.readyState === 4 && http.status === 200) {
                        spinner.style.visibility = "hidden";
                    }
                }
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