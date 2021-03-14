package main

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)



// HAY QUE DEFINIR EL TIPO DE DATO QUE VAMOS A RECIBIR EN ESTE CASO UNA PERSONA

type Persona struct{
	Id string `json:"id,omitempty"` //para ser mas explicito de lo que espero recibir
	Nombre string `json:"nombre,omitempty"` // nombre como lo espero ,  caracteristica no nulo
	Edad	int 	`json:"edad,omitempty"`
	Dir * Direccion `json:"dir,omitempty"`//que apunte al tipo de dato direccion
}

type Direccion struct{
	Zona string `json:"zona,omitempty"`
}

// AHORA PARA SIMULAR UNA BASE DE DATOS USAMOS UN JSON 
var personas []Persona



func GetPersonas(res http.ResponseWriter , req * http.Request){ // como es una rest , siempre responde y recibe , es decir escribe una respuesta y recibe una peticion
	json.NewEncoder(res).Encode(personas) // devuelve todas las personas del objeto json
}

// en este caso se pone w pero es lo mismo es con el objeto con el que escribo la respuesta
func getOnlyOnePerson(w http.ResponseWriter , req * http.Request){
	parametros := mux.Vars(req)// quine me trae todas las cosas es req 
	for _, item := range personas{
		if item.Id == parametros["id"]{
			json.NewEncoder(w).Encode(item) 
			return 
		} 
	}

	json.NewEncoder(w).Encode(&Persona{})
}

// en go las funciones en mayusculas se pueden exportar 
func CreateP(w http.ResponseWriter , req * http.Request){

	var p Persona
	_ = json.NewDecoder(req.Body).Decode(&p)// decodifico como el objeto persona
	personas = append(personas , p)
	
}


func UpdateP(res http.ResponseWriter , req * http.Request){

}

func DeleteP(w http.ResponseWriter , req * http.Request){

}



// para instalar gorilla/mux  corremos el comando > go get github.com/gorilla/mux
func main(){
	// endpoint o rutas
	router := mux.NewRouter()
	
	//cargando al inicio el json 
	personas= append(personas , Persona{Id: "1" , Nombre:"pablo" , Edad: 22 , Dir: &Direccion{Zona:"zona 4"}})


	// handlers 
	router.HandleFunc("/personas", GetPersonas).Methods("GET") // le indico que cuando se haga un get a esta ruta , realice la funcion GetPersonas
	// es un get pero con un parametro varibale llamado id
	router.HandleFunc("/personas/{id}", getOnlyOnePerson).Methods("GET")
	router.HandleFunc("/personas" , CreateP).Methods("POST") // metodos de peticion http
	router.HandleFunc("/personas/{id}" , UpdateP).Methods("POST") // metodos de peticion http
	router.HandleFunc("/personas/{id}" , DeleteP).Methods("POST") // metodos de peticion http

	log.Fatal(http.ListenAndServe(":3000",router))// le paso el puerto,enrutador
}