const loginForm = document.getElementById("login-form");
const phone = document.getElementById("phone");
const password = document.getElementById("password");
const button = loginForm.querySelector("button[type='submit']");

// Create message container for client-side validation errors
const message = document.createElement("p");
message.className = "message";
message.setAttribute("role", "alert");
// Insert after the h2 so it appears near the top of the form
const heading = loginForm.querySelector("h2");
if (heading && heading.nextSibling) {
  loginForm.insertBefore(message, heading.nextSibling);
} else {
  loginForm.appendChild(message);
}

loginForm.addEventListener("submit", function (event) {
  message.textContent = "";
  message.style.color = "";

  if (phone.value.trim() === "") {
    event.preventDefault();
    message.textContent = "Phone number is required.";
    message.style.color = "#dc2626";
    phone.focus();
    return;
  }

  if (password.value.length < 8) {
    event.preventDefault();
    message.textContent = "Password must be at least 8 characters.";
    message.style.color = "#dc2626";
    password.focus();
    return;
  }

  // Validation passed — let the form submit for real to POST /login.
  button.disabled = true;
  button.textContent = "Signing In...";
});
