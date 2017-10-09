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

var warning = {
    open: function() {
        warningFade.style.display = "block";
        document.getElementById("warningContent").textContent = "Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye Bye! Bye Bye Bye";
        //document.getElementById("warningContent").textContent = "Bye Bye";        
      
    }
}