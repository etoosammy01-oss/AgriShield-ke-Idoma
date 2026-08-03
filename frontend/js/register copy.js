document.addEventListener("DOMContentLoaded", function () {
  const registerForm = document.getElementById("register-form");

  if (!registerForm) {
    return;
  }

  const firstName = document.getElementById("first-name");
  const lastName = document.getElementById("last-name");
  const email = document.getElementById("email");
  const phone = document.getElementById("phone");
  const location = document.getElementById("location");
  const password = document.getElementById("password");
  const confirmPassword = document.getElementById("confirmPassword");
  const strengthEl = document.getElementById("password-strength");

  if (!password || !confirmPassword) {
    console.error("Register form password fields were not found.");
    return;
  }

  const button = registerForm.querySelector("button[type='submit']");

  // Message for validation errors
  const message = document.createElement("p");
  message.className = "message";
  message.setAttribute("role", "alert");
  const heading = registerForm.querySelector("h2");
  if (heading && heading.nextSibling) {
    registerForm.insertBefore(message, heading.nextSibling);
  } else {
    registerForm.appendChild(message);
  }

  // Password strength indicator
  password.addEventListener("input", function () {
    const value = password.value;
    let text = "";
    let color = "";

    if (value.length === 0) {
      text = "";
    } else if (value.length < 8) {
      text = "Weak — use at least 8 characters";
      color = "#dc2626";
    } else if (value.length < 12) {
      text = "Medium strength";
      color = "#ea580c";
    } else {
      text = "Strong password";
      color = "#16a34a";
    }

    if (strengthEl) {
      strengthEl.textContent = text;
      strengthEl.style.color = color;
    }
  });

  registerForm.addEventListener("submit", function (event) {
    message.textContent = "";
    message.style.color = "";

    if (
      firstName.value.trim() === "" ||
      lastName.value.trim() === "" ||
      email.value.trim() === "" ||
      phone.value.trim() === "" ||
      (location && location.value.trim() === "")
    ) {
      event.preventDefault();
      message.textContent = "Please complete all fields.";
      message.style.color = "#dc2626";
      return;
    }

    if (password.value !== confirmPassword.value) {
      event.preventDefault();
      message.textContent = "Passwords do not match.";
      message.style.color = "#dc2626";
      confirmPassword.focus();
      return;
    }

    if (password.value.length < 8) {
      event.preventDefault();
      message.textContent = "Password must be at least 8 characters.";
      message.style.color = "#dc2626";
      password.focus();
      return;
    }

    // Validation passed — let the form submit for real to POST /register.
    button.disabled = true;
    button.textContent = "Creating Account...";
  });
});
