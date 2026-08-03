/**
 * Modern password visibility toggle
 * Works on both login (single password) and register (password + confirm)
 * Uses eye / eye-off icons for better UX
 */
document.addEventListener("DOMContentLoaded", function () {
  const toggles = document.querySelectorAll("[data-password-toggle]");

  toggles.forEach(function (btn) {
    const targetId = btn.getAttribute("data-password-toggle");
    const input = document.getElementById(targetId);

    if (!input) return;

    btn.addEventListener("click", function () {
      const isHidden = input.type === "password";
      input.type = isHidden ? "text" : "password";

      // Update icon and accessibility
      btn.setAttribute("aria-pressed", isHidden ? "true" : "false");
      btn.setAttribute(
        "aria-label",
        isHidden ? "Hide password" : "Show password"
      );

      // Swap icon visibility
      const showIcon = btn.querySelector(".icon-show");
      const hideIcon = btn.querySelector(".icon-hide");

      if (showIcon && hideIcon) {
        showIcon.hidden = isHidden;
        hideIcon.hidden = !isHidden;
      }
    });
  });
});
