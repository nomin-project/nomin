var warningWindow = document.getElementById("warningWindow");
var warningFade = document.getElementById("warningFade");
var warningClose = document.getElementById("warningClose");

warningClose.onclick = function() {
    warningFade.style.display = "none";
}

window.onclick = function(event) {
    if (event.target == warningFade) {
        warningFade.style.display = "none";
    }
}
