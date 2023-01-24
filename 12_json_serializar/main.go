package main

import (
	"encoding/json"
	"fmt"
	"github.com/zero20/12_json_serializer/models"
)

/*
- Serialización = Codificación
función Marshal (v interfaz) ([] byte, error)

- Deserialización = Decodificación
función Unmarshal (data[] byte, v interface) error
*/
func main() {
	Deserializacion()
}

func Serializacion() {
	user := models.User{
		Id:     "1",
		Name:   "Edgar",
		Phone:  9932307383,
		Status: true,
	}
	fmt.Println(user)

	// serializar
	data, error := json.Marshal(user)

	if error != nil {
		fmt.Println("Ocurrió un error:", error)
	}

	fmt.Println(string(data))
}

func Deserializacion() {
	cadenaJson := `{"id":"1","name":"Edgar","phone":9932307383,"status":true}`
	fmt.Println(cadenaJson)
	var user models.User
	error := json.Unmarshal([]byte(cadenaJson), &user)
	if error != nil {
		fmt.Println("Error: ", error)
	}
	fmt.Println(user)
}
