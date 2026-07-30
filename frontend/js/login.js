const loginForm = document.getElementById("login-form");
const phone = document.getElementById("phone");
const password = document.getElementById("password");
const button = loginForm.querySelector("button[type='submit']");

// Create message container for client-side validation errors
const message = document.createElement("p");
message.className = "message";
loginForm.appendChild(message);

// Show/hide password toggle
const toggle = document.createElement("span");
toggle.textContent = "Show";
toggle.className = "toggle-password";
password.parentNode.appendChild(toggle);

toggle.addEventListener("click", () => {
    if (password.type === "password") {
        password.type = "text";
        toggle.textContent = "Hide";
    } else {
        password.type = "password";
        toggle.textContent = "Show";
    }
});

loginForm.addEventListener("submit", function (event) {
    message.textContent = "";

    if (phone.value.trim() === "") {
        event.preventDefault();
        message.textContent = "Phone number is required.";
        message.style.color = "red";
        phone.focus();
        return;
    }

    if (password.value.length < 8) {
        event.preventDefault();
        message.textContent = "Password must be at least 8 characters.";
        message.style.color = "red";
        password.focus();
        return;
    }

    // Validation passed — let the form submit for real to POST /login.
    // The server checks the password and creates the session.
    button.disabled = true;
    button.textContent = "Signing In...";
});
