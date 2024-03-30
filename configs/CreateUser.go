package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"main.go/DataBase"
	models "main.go/Models"
)

//TODO Trocar o ID por UUID
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha no Body"))
		return
	}
	var user models.User
	if erro = json.Unmarshal(bodyReq, &user); erro != nil {
		w.Write([]byte("Erro ao converter em struct"))
		return
	}

	db, erro := DataBase.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao Conectar no DB"))
	}
	defer db.Close()
	//TODO Criar Criptografia para armazenar senha no banco
	statement, erro := db.Prepare("insert into users (name, email, password) values (?, ?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o Statement"))
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email, user.Password)
	if erro != nil {
		w.Write([]byte("Erro ao criar o insert"))
	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao trazer o ID"))
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInsert)))
}
