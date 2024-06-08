// apiConfig
const api_cars = "http://localhost:8080/Frontend/Cars/cars.html/cars";
const api_marca = "http://localhost:8080/Frontend/Cars/cars.html/cars/marca";
const api_transmision = "http://localhost:8080/Frontend/Cars/cars.html/cars/transmision/";
const api_combustible = "http://localhost:8080/Frontend/Cars/cars.html/cars/combustible/";

// Selecciona todos los enlaces dentro de los elementos <li> del menú desplegable
let dropdownLinks = document.querySelectorAll('.dropdown-menu li a');

// Itera sobre cada enlace
dropdownLinks.forEach(link => {
    // Agrega un event listener para el evento 'click'
    link.addEventListener('click', function(event) {
        // Previene la acción predeterminada del enlace
        event.preventDefault();
        
        // Obtiene el valor del texto dentro del enlace
        let linkText = event.target.textContent;

        // Obtiene el ID del elemento al que se le dio clic
        let elementId = event.target.parentElement.parentElement.parentElement.parentElement.id;
        console.log(elementId);

        switch(elementId) {
            case 'Marca':
                api_url = api_marca;
                break;
            case 'transmision':
                api_url = api_transmision;
                break;
            case 'combustible':
                api_url = api_combustible;
                break;
            default:
                api_url = api_cars;
        }

        // Imprime el texto del enlace
        filtername = document.getElementById("name");
        filtername.innerHTML = linkText;
        url = api_url + "/" + linkText;
        console.log(url);
    });
});

//Creación de tarjetas
function cards(data){
    const cards = document.getElementById("cards");
    cards.innerHTML = "";
  
    for(i=0; i<data.length; i++){
        const dato = data[i];

        const container = document.createElement("div");
        container.className = "card";

        const img = document.createElement("img");
        img.src = dato.Imagen;

        const card = document.createElement("div");
        card.className = "card-body";

        const title = document.createElement("h5");
        title.className = "card-title";
        title.textContent = dato.Modelo;

        const p = document.createElement("p");
        p.className = "card-text";
        p.textContent = "Marca: " + dato.Marca + " | Transmisión: " + dato.Transmision + " | Combustible: " + dato.Combustible;

        const more = document.createElement("a");
        more.className = "btn btn-primary";
        more.href = "#";
        more.textContent = "Reservar";


        container.appendChild(img);
        container.appendChild(card);
        card.appendChild(title);
        card.appendChild(p);
        card.appendChild(more);

        cards.appendChild(container);
    }
  }

async function consultar_Carro(url) {
    try {
      const response = await axios.get(url);
      console.log(response);
      return response;
    } catch (error) {
      console.error(`Falló esta vuelta: ${error}`);
      alert("API query failed");
    }
  }