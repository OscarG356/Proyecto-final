CREATE TABLE IF NOT EXISTS Carros (
    id SERIAL PRIMARY KEY,
    marca VARCHAR(100),
    modelo VARCHAR(100),
    tipo VARCHAR(100),
    transmision VARCHAR(100),
    combustible VARCHAR(100),
	color VARCHAR(100),
	ubicacion VARCHAR(100),
	imagen VARCHAR(255)
);

INSERT INTO carros (marca, modelo, tipo, transmision, combustible, color, ubicacion, imagen) VALUES
   ('Toyota', 'Corolla', 'Sedan', 'Automática', 'Gasolina', 'Negro', 'Bogotá', '\Backend\images\corola_toyota.png'),
   ('Honda', 'Civic', 'Sedan', 'Manual', 'Gasolina', 'Azul', 'Medellín', '\Backend\images\honda_civic.png'),
   ('Ford', 'F-150', 'Camioneta', 'Automática', 'Gasolina', 'Rojo', 'Cali', '\Backend\images\ford_f150.jpeg'),
   ('Chevrolet', 'Camaro', 'Deportivo', 'Automática', 'Gasolina', 'Amarillo', 'Barranquilla', '\Backend\images\camaro_chevrolet.png'),
   ('Volkswagen', 'Golf', 'Hatchback', 'Manual', 'Diesel', 'Blanco', 'Cartagena', '\Backend\images\golf_volks.png');

CREATE TABLE IF NOT EXISTS Usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    contraseña VARCHAR(100),
	reserva VARCHAR(20)
);

INSERT INTO usuarios (nombre, contraseña, telefono, reserva) VALUES
    ('Admin', '1234', '123456789', '1');
