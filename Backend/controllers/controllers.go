package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/OscarG356/Proyecto_final-Desarrollo_Web/Backend/models"
	"github.com/OscarG356/Proyecto_final-Desarrollo_Web/Backend/repository"
)

var (
	ListQueryC   = "SELECT * FROM Carro limit $1 offset $2"
	ReadQueryC   = "SELECT * FROM Carro WHERE id=$1"
	CreateQueryC = "INSERT INTO Carro (marca, referencia, modelo, tipo, potencia, torque, transmision, motor, pasajeros, combustible, consumo, almacenamiento, descripcion, lujo, deportivo, imagen, reservado) VALUES (:marca, :referencia, :modelo, :tipo, :potencia, :torque, :transmision, :motor, :pasajeros, :combustible, :consumo, :almacenamiento, :descripcion, :lujo, :deportivo, :imagen, :reservado) RETURNING id"
	UpdateQueryC = "UPDATE Carro SET %s WHERE id=:id"
	DeleteQueryC = "DELETE FROM Carro WHERE id=$1"
)
var (
	ListQueryU   = "SELECT * FROM Usuario limit $1 offset $2"
	ReadQueryU   = "SELECT * FROM Usuario WHERE id=$1"
	CreateQueryU = "INSERT INTO Usuario (nombre, contrasena, reservas) VALUES (:nombre, :contrasena, :reservas) RETURNING id"
	UpdateQueryU = "UPDATE Usuario SET %s WHERE id=:id"
	DeleteQueryU = "DELETE FROM Usuario WHERE id=$1"
)

type Controller struct {
	repo repository.Repository[models.Carro]
}

type ControllerU struct {
	repo repository.Repository[models.Usuario]
}

func Newcontroller(repo repository.Repository[models.Carro]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("se requiere un repositorio válido para un controlador")
	}
	return &Controller{
		repo: repo,
	}, nil
}
func NewcontrollerU(repo repository.Repository[models.Usuario]) (*ControllerU, error) {
	if repo == nil {
		return nil, fmt.Errorf("se requiere un repositorio válido para un controlador")
	}
	return &ControllerU{
		repo: repo,
	}, nil
}

func (c *Controller) ListarCarros(limit, offset int) ([]byte, error) {
	carros, _, err := c.repo.List(context.TODO(), ListQueryC, limit, offset)
	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}

	jsonCarro, err := json.Marshal(carros)

	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}
	return jsonCarro, nil
}

func (c *Controller) ListarCarrosMarca(marca string, limit, offset int) ([]byte, error) {
	carros, _, err := c.repo.List(context.TODO(), ListQueryC, limit, offset)
	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}

	var carromarca []*models.Carro

	for _, carro := range carros {
		if carro.Marca == marca {
			carromarca = append(carromarca, carro)
		}
	}

	jsonCarro, err := json.Marshal(carromarca)

	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}
	return jsonCarro, nil
}

func (c *Controller) ListarCarrosTipo(tipo string, limit, offset int) ([]byte, error) {
	carros, _, err := c.repo.List(context.TODO(), ListQueryC, limit, offset)
	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}

	var carrotipo []*models.Carro

	for _, carro := range carros {
		if carro.Tipo == tipo {
			carrotipo = append(carrotipo, carro)
		}
	}

	jsonCarro, err := json.Marshal(carrotipo)

	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}
	return jsonCarro, nil
}
func (c *Controller) ListarCarrosTransmision(transmision string, limit, offset int) ([]byte, error) {
	carros, _, err := c.repo.List(context.TODO(), ListQueryC, limit, offset)
	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}

	var carrotransmision []*models.Carro

	for _, carro := range carros {
		if carro.Transmision == transmision {
			carrotransmision = append(carrotransmision, carro)
		}
	}

	jsonCarro, err := json.Marshal(carrotransmision)

	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}

	return jsonCarro, nil
}

func (c *Controller) ListarCarrosCombustible(combustible string, limit, offset int) ([]byte, error) {
	carros, _, err := c.repo.List(context.TODO(), ListQueryC, limit, offset)
	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}

	var carrocombustible []*models.Carro

	for _, carro := range carros {
		if carro.Combustible == combustible {
			carrocombustible = append(carrocombustible, carro)
		}
	}

	jsonCarro, err := json.Marshal(carrocombustible)

	if err != nil {
		log.Printf("falló al obtener los carros, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener los carros, error: %s", err.Error())
	}
	return jsonCarro, nil
}

func (c *ControllerU) ListarUsuarios(limit, offset int) ([]byte, error) {
	usuarios, _, err := c.repo.List(context.TODO(), ListQueryU, limit, offset)
	if err != nil {
		log.Printf("fallo al obtener usuarios, error: %s", err.Error())
		return nil, fmt.Errorf("fallo al obtener usuarios, error: %s", err.Error())
	}

	jsonUsuarios, err := json.Marshal(usuarios)

	if err != nil {
		log.Printf("fallo al obtener usuarios, error: %s", err.Error())
		return nil, fmt.Errorf("fallo al obtener usuarios, error: %s", err.Error())
	}
	return jsonUsuarios, nil
}

func (c *Controller) TraerCarro(id string) ([]byte, error) {
	carro, err := c.repo.Read(context.TODO(), ReadQueryC, id)
	if err != nil {
		log.Printf("falló al obtener un carro, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener un carro, error: %s", err.Error())
	}
	Carrojson, err := json.Marshal(carro)
	if err != nil {
		log.Printf("falló al obtener un carro, error: %s", err.Error())
		return nil, fmt.Errorf("falló al obtener un carro, error: %s", err.Error())
	}
	return Carrojson, nil
}

func (c *ControllerU) TraerUsuario(id string) ([]byte, error) {
	usuario, err := c.repo.Read(context.TODO(), ReadQueryU, id)
	if err != nil {
		log.Printf("fallo al obtener un usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al obtener un usuario, con error: %s", err.Error())
	}
	Usuariojson, err := json.Marshal(usuario)
	if err != nil {
		log.Printf("fallo al obtener un usuario, error: %s", err.Error())
		return nil, fmt.Errorf("fallo al obtener un usuario, error: %s", err.Error())
	}
	return Usuariojson, nil
}

func (c *Controller) CrearCarro(body []byte) (int64, error) {
	NuevoCarro := &models.Carro{}
	err := json.Unmarshal(body, NuevoCarro)
	if err != nil {
		log.Printf("fallo al crear un carro,  error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un carro,  error: %s", err.Error())
	}
	valores_columnas := map[string]any{
		"id":             NuevoCarro.Id,
		"marca":          NuevoCarro.Marca,
		"referencia":     NuevoCarro.Referencia,
		"modelo":         NuevoCarro.Modelo,
		"tipo":           NuevoCarro.Tipo,
		"potencia":       NuevoCarro.Potencia,
		"torque":         NuevoCarro.Torque,
		"transmision":    NuevoCarro.Transmision,
		"motor":          NuevoCarro.Motor,
		"pasajeros":      NuevoCarro.Pasajeros,
		"combustible":    NuevoCarro.Combustible,
		"consumo":        NuevoCarro.Consumo,
		"almacenamiento": NuevoCarro.Almacenamiento,
		"descripcion":    NuevoCarro.Descripcion,
		"lujo":           NuevoCarro.Lujo,
		"deportivo":      NuevoCarro.Deportivo,
		"imagen":         NuevoCarro.Imagen,
		"reservado":      NuevoCarro.Reservado,
	}
	NuevoId, err := c.repo.Create(context.TODO(), CreateQueryC, valores_columnas)
	if err != nil {
		log.Printf("fallo al crear un carro,  error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un carro,  error: %s", err.Error())
	}
	return NuevoId, nil
}

func (c *ControllerU) CrearUsuario(body []byte) (int64, error) {
	NuevoUsuario := &models.Usuario{}
	err := json.Unmarshal(body, NuevoUsuario)
	if err != nil {
		log.Printf("fallo al crear un usuario, error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un usuario, error: %s", err.Error())
	}
	valores_columnas := map[string]any{
		"id":         NuevoUsuario.Id,
		"nombre":     NuevoUsuario.Nombre,
		"contrasena": NuevoUsuario.Contrasena,
		"reservas":   NuevoUsuario.Reservas,
	}
	NuevoId, err := c.repo.Create(context.TODO(), CreateQueryU, valores_columnas)
	if err != nil {
		log.Printf("fallo al crear un usuario, error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un usuario, error: %s", err.Error())
	}
	return NuevoId, nil
}

func buildUpdateQuery(columnasActualizar map[string]any) string {
	columnas := []string{}
	for key := range columnasActualizar {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return columnasString
}

func (c *Controller) ActualizarCarro(body []byte, id string) error {
	valoresActualizarBody := make(map[string]any)
	err := json.Unmarshal(body, &valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un carro, error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un carro, error: %s", err.Error())
	}
	if len(valoresActualizarBody) == 0 {
		log.Printf("fallo al actualizar un carro, error: no hay datos en el body")
		return fmt.Errorf("fallo al actualizar un carro, error:no hay datos en el body")
	}
	UpdtQueryC := fmt.Sprintf(UpdateQueryC, buildUpdateQuery(valoresActualizarBody))
	valoresActualizarBody["id"] = id
	err = c.repo.Update(context.TODO(), UpdtQueryC, valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un carro, error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un carro, error: %s", err.Error())
	}
	return nil
}

func (c *ControllerU) ActualizarUsuario(body []byte, id string) error {
	valoresActualizarBody := make(map[string]any)
	err := json.Unmarshal(body, &valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, error: %s", err.Error())
	}
	if len(valoresActualizarBody) == 0 {
		log.Printf("fallo al actualizar un usuario, error: no hay datos en el body")
		return fmt.Errorf("fallo al actualizar un usuario, error:no hay datos en el body")
	}
	UpdtQueryU := fmt.Sprintf(UpdateQueryU, buildUpdateQuery(valoresActualizarBody))
	valoresActualizarBody["id"] = id
	err = c.repo.Update(context.TODO(), UpdtQueryU, valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, error: %s", err.Error())
	}
	return nil
}

func (c *Controller) EliminarCarro(id string) error {
	err := c.repo.Delete(context.TODO(), DeleteQueryC, id)
	if err != nil {
		log.Printf("fallo al eliminar un carro, error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un carro, error: %s", err.Error())
	}
	return nil
}

func (c *ControllerU) EliminarUsuario(id string) error {
	err := c.repo.Delete(context.TODO(), DeleteQueryU, id)
	if err != nil {
		log.Printf("fallo al eliminar un usuario, error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un usuario, error: %s", err.Error())
	}
	return nil
}
