document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('contactForm');
    const thanksMessage = document.getElementById('thanksMessage');

    form.addEventListener('submit', function (event) {
        event.preventDefault();

        const formData = new FormData(form);

        fetch(form.action, {
            method: 'POST',
            body: formData
        })
        .then(response => {
            if (response.ok) {
                form.reset();
                thanksMessage.style.display = 'flex';
            } else {
                alert('There was an error. Please try again.');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('There was an error. Please try again.');
        });
    });
});
