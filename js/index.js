
function updateBadgeColor() {
    windValObj = document.getElementById("windVal")
    windValue = windValObj.innerText
    windStatObj = document.getElementById("windStatus")

    if (windValue <= 6) {
        windStatObj.className="badge text-bg-light"
    } else if ((windValue >= 7) && (windValue <= 15)) {
        windStatObj.className="badge text-bg-warning"
    } else if (windValue > 15){
        windStatObj.className="badge text-bg-danger"
    }

    waterValObj = document.getElementById("waterVal")
    waterValue = waterValObj.innerText
    waterStatObj = document.getElementById("waterStatus")


    if (waterValue < 5) {
        waterStatObj.className="badge text-bg-light"
    } else if ((waterValue >= 6) && (waterValue <= 8)) {
        waterStatObj.className="badge text-bg-warning"
    } else if (waterValue >= 8){
        waterStatObj.className="badge text-bg-danger"
    }
}