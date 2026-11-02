const username = document.getElementById("username-input");
const password = document.getElementById("password-input");
const button = document.getElementById("button");
const result = document.getElementById("result");

button.addEventListener("click", async function (){
    const response = await fetch("/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: username.value,
            password: password.value
        }),
    })

    if (!response.ok) {
        console.log("Ошибка входа: ", response.status);
        result.textContent = "Ошибка входа";
    } else {
        console.log("Успешный вход!");
        result.textContent = "Успешный вход!";
        document.location = "/home.html"
    }
});
