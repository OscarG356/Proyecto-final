package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/OscarG356/Proyecto_final-Desarrollo_Web/Backend/controllers"
	"github.com/gorilla/mux"
)

type HandlerCarros struct {
	controller *controllers.Controller
}

type HandlerUsuarios struct {
	controller *controllers.ControllerU
}

func NewHandlerCarros(controller *controllers.Controller) (*HandlerCarros, error) {
	if controller == nil {
		return nil, fmt.Errorf("se requiere un controlador no nulo para instanciar un handler")
	}
	return &HandlerCarros{
		controller: controller,
	}, nil
}

func NewHandlerUsuarios(controller *controllers.ControllerU) (*HandlerUsuarios, error) {
	if controller == nil {
		return nil, fmt.Errorf("se requiere un controlador no nulo para instanciar un handler")
	}
	return &HandlerUsuarios{
		controller: controller,
	}, nil
}

func (hc *HandlerCarros) ListarCarros(w http.ResponseWriter, r *http.Request) {
	carros, err := hc.controller.ListarCarros(100, 0)
	if err != nil {
		log.Printf("fallo al leer carros, error: %s", err.Error())
		http.Error(w, "fallo al leer carros", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carros)
}
func (hc *HandlerCarros) ListarCarrosMarca(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	marca := vars["marca"]

	carros, err := hc.controller.ListarCarrosMarca(marca, 100, 0)
	if err != nil {
		log.Printf("fallo al leer carros por marca, error: %s", err.Error())
		http.Error(w, "fallo al leer carros por marca", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carros)
}
func (hc *HandlerCarros) ListarCarrosTipo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tipo := vars["tipo"]

	carros, err := hc.controller.ListarCarrosTipo(tipo, 100, 0)
	if err != nil {
		log.Printf("fallo al leer carros por tipo, error: %s", err.Error())
		http.Error(w, "fallo al leer carros por tipo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carros)
}

func (hc *HandlerCarros) ListarCarrosTransmision(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	transmision := vars["transmision"]

	carros, err := hc.controller.ListarCarrosTransmision(transmision, 100, 0)
	if err != nil {
		log.Printf("fallo al leer carros por transmision, error: %s", err.Error())
		http.Error(w, "fallo al leer carros por transmision", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carros)
}

func (hc *HandlerCarros) ListarCarrosCombustible(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	combustible := vars["combustible"]

	carros, err := hc.controller.ListarCarrosCombustible(combustible, 100, 0)
	if err != nil {
		log.Printf("fallo al leer carros por combustible, error: %s", err.Error())
		http.Error(w, "fallo al leer carros por combustible", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carros)
}

func (hc *HandlerUsuarios) ListarUsuarios(w http.ResponseWriter, r *http.Request) {

	usuarios, err := hc.controller.ListarUsuarios(100, 0)
	if err != nil {
		log.Printf("fallo al leer usuarios, error: %s", err.Error())
		http.Error(w, "fallo al leer usuarios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usuarios)
}

func (hc *HandlerCarros) CrearCarros(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al crear nuevo carro, error: %s", err.Error())
		http.Error(w, "fallo al crear nuevo carro", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	nuevoId, err := hc.controller.CrearCarro(body)
	if err != nil {
		log.Printf("fallo al crear nuevo carro, error: %s", err.Error())
		http.Error(w, "fallo al crear nuevo carro", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, fmt.Sprintf("id del nuevo carro: %d", nuevoId))
}
func (hc *HandlerUsuarios) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al crear nuevo usuario, error: %s", err.Error())
		http.Error(w, "fallo al crear nuevo usuario", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	nuevoId, err := hc.controller.CrearUsuario(body)
	if err != nil {
		log.Printf("fallo al crear nuevo usuario, error: %s", err.Error())
		http.Error(w, "fallo al crear nuevo usuario", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, fmt.Sprintf("id del nuevo usuario: %d", nuevoId))
}

func (hc *HandlerCarros) TraerCarro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	carro, err := hc.controller.TraerCarro(id)
	if err != nil {
		log.Printf("fallo a leer un carro, error: %s", err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("el carro con id %s no se pudo encontrar", id)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(carro)
}

func (hc *HandlerUsuarios) TraerUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	usuario, err := hc.controller.TraerUsuario(id)
	if err != nil {
		log.Printf("fallo a leer un usuario, error: %s", err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("el usuario con id %s no se pudo encontrar", id)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usuario)
}

func (hc *HandlerCarros) ActualizarCarro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al actualizar un carro, error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al actualizar un carro, error: %s", err.Error()), http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = hc.controller.ActualizarCarro(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un carro, error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al actualizar un carro, error: %s", err.Error()), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (hc *HandlerUsuarios) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al actualizar un usuario, error: %s", err.Error()), http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = hc.controller.ActualizarUsuario(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al actualizar un usuario, error: %s", err.Error()), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (hc *HandlerCarros) EliminarCarro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := hc.controller.EliminarCarro(id)
	if err != nil {
		log.Printf("fallo al Eliminar un carro, error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al eliminar un carro, error: %s", err.Error()), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (hc *HandlerUsuarios) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := hc.controller.EliminarUsuario(id)
	if err != nil {
		log.Printf("fallo al Eliminar un usuario, con error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al eliminar un usuario, error: %s", err.Error()), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
