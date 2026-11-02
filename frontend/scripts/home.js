
async function Get_data() {
    try {

        const text = document.getElementById("data");

        const response = await fetch(`https://api.demonlist.org/level/classic/get?placement=1`);
        const result = await response.json();

        const levelName = result.data.name;

        text.textContent = "Name: " + levelName + "Placement: 1";

    } catch (error) {
        console.error("Ошибка при запросе:", error);
    }
}

Get_data();
