const usernameInput = document.getElementById("username-input");
const passwordInput = document.getElementById("password-input");
const button = document.getElementById("username-button");
const result = document.getElementById("result");

button.addEventListener("click", async function () {
    const response = await fetch("/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: usernameInput.value,
            password: passwordInput.value,
        }),
    });

    if (!response.ok) {
        console.error("Ошибка:", response.status);
        result.textContent = "Ошибка";
    } else {
        console.log("Зарегистрировано");
        result.textContent = "Успешно";
        document.location = "/login.html"
    }
});
