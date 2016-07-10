function check() {
    var url = document.getElementById("url").value;
    
    if (url.substr(0, 7) === "http://" || url.substr(0, 8) === "https://") {
        return true;
    } else {
        document.getElementById("url").value = "http://" + url;
        return true;
    }
}