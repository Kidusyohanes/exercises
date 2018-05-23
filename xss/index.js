"use strict";

const name = document.getElementById("name");
const nameOut = document.getElementById("name-out");

document.getElementById("name-form").addEventListener("submit", evt => {
    evt.preventDefault();
    nameOut.innerHTML = name.value;
    //nameOut.textContent = name.value;
});


