CREATE TABLE IF NOT EXISTS Carro (
    id SERIAL PRIMARY KEY,
    marca VARCHAR(100),
    modelo VARCHAR(100),
    tipo VARCHAR(100),
    transmision VARCHAR(100),
    combustible VARCHAR(100),
	color VARCHAR(100),
	ubicacion VARCHAR(100),
    reservado BOOLEAN,
	imagen VARCHAR(255)
);

INSERT INTO Carro (marca, modelo, tipo, transmision, combustible, color, ubicacion, reservado, imagen) VALUES
   ('Toyota', 'Corolla', 'Sedan', 'Automática', 'Gasolina', 'Negro', 'Bogotá', true, '/Backend/images/corola_toyota.png'),
   ('Honda', 'Civic', 'Sedan', 'Manual', 'Gasolina', 'Azul', 'Medellín', true, '/Backend/images/honda_civic.png'),
   ('Ford', 'F-150', 'Camioneta', 'Automática', 'Gasolina', 'Rojo', 'Cali', true, '/Backend/images/ford_f150.jpeg'),
   ('Chevrolet', 'Camaro', 'Deportivo', 'Automática', 'Gasolina', 'Amarillo', true, '/Backend/images/camaro_chevrolet.png'),
   ('Volkswagen', 'Golf', 'Hatchback', 'Manual', 'Diesel', 'Blanco', 'Cartagena', false, '/Backend/images/golf_volks.png');

CREATE TABLE IF NOT EXISTS Usuario (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    contrasena VARCHAR(100),
    reservas VARCHAR(20)
);

INSERT INTO Usuario (nombre, contrasena, reservas) VALUES
    ('Admin', '1234', '1');
