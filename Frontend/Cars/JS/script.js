console.log("Holi");

//********************************************** */
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

        // Imprime el texto del enlace
        console.log(linkText);
    });
});
//********************************************** */