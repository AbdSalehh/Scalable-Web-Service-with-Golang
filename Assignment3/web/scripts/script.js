const updateStatus = (status) => {
    let water = status.water;
    let wind = status.wind;

    let waterStatus, windStatus, waterDesc, windDesc, waterClass, windClass;

    if (water < 5) {
        waterStatus = 'Aman';
        waterDesc = 'Pada keadaan yang sangat aman';
        waterClass = 'air-aman';
    } else if (water >= 6 && water <= 8) {
        waterStatus = 'Siaga';
        waterDesc = 'Berpotensi dalam keadaan siaga';
        waterClass = 'air-siaga';
    } else {
        waterStatus = 'Bahaya';
        waterDesc = 'Dalam keadaan yang berbahaya';
        waterClass = 'air-bahaya';
    }

    if (wind < 6) {
        windStatus = 'Aman';
        windDesc = 'Cuaca sangat aman';
        windClass = 'angin-aman';
    } else if (wind >= 7 && wind <= 15) {
        windStatus = 'Siaga';
        windDesc = 'Cuaca dalam keadaan siaga';
        windClass = 'angin-siaga';
    } else {
        windStatus = 'Bahaya';
        windDesc = 'Angin sangat kencang';
        windClass = 'angin-bahaya';
    }

    document.getElementById('waterNumber').innerText = water;
    document.getElementById('waterStatus').innerText = waterStatus;
    document.getElementById('waterDesc').innerText = waterDesc;
    document.getElementById('waterImage').className = waterClass;

    document.getElementById('windNumber').innerText = wind;
    document.getElementById('windStatus').innerText = windStatus;
    document.getElementById('windDesc').innerText = windDesc;
    document.getElementById('windImage').className = windClass;
};

const fetchStatusData = () => {
    fetch('/status')
        .then((response) => response.json())
        .then((data) => {
            updateStatus(data.status);
        })
        .catch((error) => console.error('Error fetching status:', error));
};

fetchStatusData();

setInterval(fetchStatusData, 15000);
