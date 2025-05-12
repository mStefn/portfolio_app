document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("contactForm");
    const thanksMessage = document.getElementById("thanksMessage");

    if (form) {
        form.addEventListener("submit", function (e) {
            e.preventDefault();

            fetch(form.action, {
                method: "POST",
                body: new FormData(form),
                headers: {
                    'Accept': 'application/json'
                }
            }).then(response => {
                if (response.ok) {
                    form.style.display = "none";
                    thanksMessage.style.display = "flex"; 
                } else {
                    alert("There was a problem sending the message. Please try again.");
                }
            }).catch(() => {
                alert("There was a problem connecting to the server.");
            });
        });
    }
});
