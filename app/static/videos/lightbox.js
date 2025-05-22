document.addEventListener("DOMContentLoaded", () => {
    const lightbox = document.getElementById("lightbox");
    const lightboxImg = document.querySelector(".lightbox-img");
    const closeBtn = document.querySelector(".lightbox-close");

    const triggers = document.querySelectorAll(".lightbox-trigger");

    triggers.forEach(img => {
        img.addEventListener("click", () => {
            lightboxImg.src = img.src;
            lightbox.classList.add("active");
        });
    });

    closeBtn.addEventListener("click", () => {
        lightbox.classList.remove("active");
        lightboxImg.src = "";
    });

    lightbox.addEventListener("click", (e) => {
        if (e.target === lightbox) {
            lightbox.classList.remove("active");
            lightboxImg.src = "";
        }
    });

    document.addEventListener("keydown", (e) => {
        if (e.key === "Escape" && lightbox.classList.contains("active")) {
            lightbox.classList.remove("active");
            lightboxImg.src = "";
        }
    });
});
