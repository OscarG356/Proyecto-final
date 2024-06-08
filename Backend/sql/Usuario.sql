CREATE TABLE IF NOT EXISTS Usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    contraseña VARCHAR(100),
	reserva VARCHAR(20)
);