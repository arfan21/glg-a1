const init = async () => {
    const waterElm = document.getElementById("water");
    const windElm = document.getElementById("wind");
    const result = await fetch("/weather-status.json", { method: "GET" });
    const response = await result.json();

    let colorWater = "";
    let colorWind = "";

    if (response.status.water < 5) {
        colorWater = "green";
    }

    if (response.status.water >= 6 && response.status.water <= 8) {
        colorWater = "yellow";
    }

    if (response.status.water > 5) {
        colorWater = "red";
    }

    if (response.status.wind < 6) {
        colorWind = "green";
    }

    if (response.status.wind >= 7 && response.status.wind <= 15) {
        colorWind = "yellow";
    }

    if (response.status.wind > 15) {
        colorWind = "red";
    }

    waterElm.innerHTML = `
        <p class="title">Water</p>
        <h1 class="status-value" style="color: ${colorWater};">${response.status.water}</h1>
    `;

    windElm.innerHTML = `
        <p class="title">Wind</p>
        <h1 class="status-value" style="color: ${colorWind};">${response.status.wind}</h1>
    `;

    setInterval(async () => {
        const result = await fetch("/weather-status.json", { method: "GET" });
        const response = await result.json();
        let colorWater = "";
        let colorWind = "";

        if (response.status.water < 5) {
            colorWater = "green";
        }

        if (response.status.water >= 6 && response.status.water <= 8) {
            colorWater = "yellow";
        }

        if (response.status.water > 5) {
            colorWater = "red";
        }

        if (response.status.wind < 6) {
            colorWind = "green";
        }

        if (response.status.wind >= 7 && response.status.wind <= 15) {
            colorWind = "yellow";
        }

        if (response.status.wind > 15) {
            colorWind = "red";
        }

        waterElm.innerHTML = `
        <p class="title">Water</p>
        <h1 class="status-value" style="color: ${colorWater};">${response.status.water}</h1>
    `;

        windElm.innerHTML = `
        <p class="title">Wind</p>
        <h1 class="status-value" style="color: ${colorWind};">${response.status.wind}</h1>
    `;
    }, 5000);
};

init();
