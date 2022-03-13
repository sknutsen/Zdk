//const btn = document.getElementById("navbar-toggler-lg");
//const brandName = document.getElementById("brand-name");
//const navbar = document.getElementById("navbar");

//btn.onclick = toggleNav;

//function toggleNav(t, ev) {
//    if (navbar.style.width === "250px" || navbar.style.width === undefined || navbar.style.width === "") {
//        closeNav();

//    } else {
//        openNav();
//    }
//}

//function openNav() {
//    navbar.style.width = "250px";
//    navbar.style.paddingLeft = "18";
//    brandName.classList.remove("sidebar-center");
//    let texts = document.getElementsByClassName("navlink-text");
//    let icons = document.getElementsByClassName("nav-item-icon");

//    for (let i = 0; i < icons.length; i++) {
//        icons[i].classList.remove("nav-item-icon-lg");
//        icons[i].classList.remove("sidebar-center");
//    }

//    setTimeout(function () {
//        for (let i = 0; i < texts.length; i++) {
//            texts[i].style.display = "flex";
//        }
//    }, 300);
//}

//function closeNav() {
//    navbar.style.width = "90px";
//    navbar.style.paddingLeft = "0 !important";
//    brandName.classList.add("sidebar-center");
//    let texts = document.getElementsByClassName("navlink-text");
//    let icons = document.getElementsByClassName("nav-item-icon");

//    for (let i = 0; i < texts.length; i++) {
//        texts[i].style.display = "none";
//    }

//    for (let i = 0; i < icons.length; i++) {
//        icons[i].classList.add("nav-item-icon-lg");
//        icons[i].classList.add("sidebar-center");
//    }
//}

function toggleNav(t, ev) {
    let navbar = document.getElementById("navbar");
    let brandName = document.getElementById("brand-name");
    let texts = document.getElementsByClassName("navlink-text");
    let icons = document.getElementsByClassName("nav-item-icon");

    if (navbar.style.width === "250px" || navbar.style.width === undefined || navbar.style.width === "") {
        closeNav(navbar, brandName, texts, icons);

    } else {
        openNav(navbar, brandName, texts, icons);
    }
}

function openNav(navbar, brandName, texts, icons) {
    navbar.style.width = "250px";
    navbar.style.paddingLeft = "18";
    brandName.classList.remove("sidebar-center");

    for (let i = 0; i < icons.length; i++) {
        icons[i].classList.remove("nav-item-icon-lg");
        icons[i].classList.remove("sidebar-center");
    }

    setTimeout(function () {
        for (let i = 0; i < texts.length; i++) {
            texts[i].style.display = "flex";
        }
    }, 300);
}

function closeNav(navbar, brandName, texts, icons) {
    navbar.style.width = "90px";
    navbar.style.paddingLeft = "0 !important";
    brandName.classList.add("sidebar-center");

    for (let i = 0; i < texts.length; i++) {
        texts[i].style.display = "none";
    }

    for (let i = 0; i < icons.length; i++) {
        icons[i].classList.add("nav-item-icon-lg");
        icons[i].classList.add("sidebar-center");
    }
}

