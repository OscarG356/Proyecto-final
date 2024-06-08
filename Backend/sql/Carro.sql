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