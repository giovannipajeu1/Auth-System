package configs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/DataBase"
	models "main.go/Models"
)

//TODO Atualziar funcao para nao trazer Id e Password
func ReadUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := DataBase.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao Conectar no Banco"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from users")
	if erro != nil {
		w.Write([]byte("Erro ao Conectar no Banco"))
		return
	}
	defer linhas.Close()

	var users []models.User
	for linhas.Next() {
		var user models.User

		if erro := linhas.Scan(&user.ID, &user.Name, &user.Email, &user.Password); erro != nil {
			w.Write([]byte("Erro ao Scan User"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro ao converter para json"))
		return
	}

}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter ID"))
		return
	}
	db, erro := DataBase.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao Conectar no Banco"))
		return
	}

	linha, erro := db.Query("select * from users where id = ? ", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar user"))
		return
	}

	var user models.User
	if linha.Next() {
		if erro := linha.Scan(&user.ID, &user.Name, &user.Email, &user.Password); erro != nil {
			w.Write([]byte("Erro ao Scan User"))
			return
		}
		w.WriteHeader(http.StatusOK)
		if erro := json.NewEncoder(w).Encode(user); erro != nil {
			w.Write([]byte("Erro ao converter para json"))
			return
		}

	}
}
