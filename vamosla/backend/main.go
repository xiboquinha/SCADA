package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

//placeholder
func TelaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tela inicial")
	fmt.Fprintln(w, "Tela inicial")
}

//plcaeholder
func ConecaIHM(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tela ihm")
	fmt.Fprintln(w, "Lugar pra conectar IHM")
}

func ConectaCLP() {
	server := NewServer(":3121")
	go server.Start()

}

func TelaCLP(w http.ResponseWriter, r *http.Request) {
	db, err := initDB()
	if err != nil {
		http.Error(w, "erro ao acessar bando de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM msgs ORDER BY timestamp DESC")
	if err != nil {
		http.Error(w, "erro ao consultar banco de dados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var msgs []struct {
		ID        int    `json:"ID"`
		Msg       string `json:"Msg"`
		TimeStamp string `json:"TimeStamp"`
	}

	for rows.Next() {
		var m struct {
			ID        int
			Msg       string
			TimeStamp string
		}
		if err := rows.Scan(&m.ID, &m.Msg, &m.TimeStamp); err != nil {
			http.Error(w, "erro lendo linha", http.StatusInternalServerError)
			return
		}
		msgs = append(msgs, struct {
			ID        int    `json:"ID"`
			Msg       string `json:"Msg"`
			TimeStamp string `json:"TimeStamp"`
		}{
			ID:        m.ID,
			Msg:       m.Msg,
			TimeStamp: m.TimeStamp,
		})
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "erro ao processar linhas do banco de dados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msgs)
}

func handlerDeleteMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		fmt.Println("delete requisitad")
		idStr := r.URL.Query().Get("id")
		fmt.Println("id recebido: ", idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("id inválido", idStr)
			http.Error(w, "ID invalido", http.StatusBadRequest)
			return
		}

		err = DeletaLinha(db, id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Falha ao excluir: ", http.StatusInternalServerError)
			return
		}

		fmt.Println("deu certoooo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("linha excluida"))
	} else {
		fmt.Println("deu merda", r.Method)
		http.Error(w, "metodo nao permitido", http.StatusMethodNotAllowed)
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dashdash")
	db, err := initDB()
	if err != nil {
		http.Error(w, "erro ao acessar banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID, Msg, TimeStamp FROM msgs ORDER BY timestamp")
	if err != nil {
		http.Error(w, "erro ao consultar banco de dados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var dadosDashboard []struct {
		ID        int    `json:"ID"`
		Msg       string `json:"Msg"`
		TimeStamp string `json:"TimeStamp"`
	}

	for rows.Next() {
		var d struct {
			ID        int
			Msg       string
			TimeStamp string
		}
		if err := rows.Scan(&d.ID, &d.Msg, &d.TimeStamp); err != nil {
			http.Error(w, "erro lendo a linha tal", http.StatusInternalServerError)
			return
		}
		dadosDashboard = append(dadosDashboard, struct {
			ID        int    `json:"ID"`
			Msg       string `json:"Msg"`
			TimeStamp string `json:"TimeStamp"`
		}{
			ID:        d.ID,
			Msg:       d.Msg,
			TimeStamp: d.TimeStamp,
		})
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "erro ao processar linhas do banco de dados ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dadosDashboard)
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	go ConectaCLP()

	mux.HandleFunc("/", TelaInicial)
	mux.HandleFunc("/ihm", ConecaIHM)
	mux.HandleFunc("/clp", TelaCLP)
	mux.HandleFunc("/clp/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tentao excluir lll")
		handlerDeleteMessage(db, w, r)
	})
	mux.HandleFunc("/dashboard", dashboard)

	corshandler := cors.Default().Handler(mux)

	fmt.Println("8080")
	if err := http.ListenAndServe(":8080", corshandler); err != nil {
		fmt.Println("ERROOO NAO COMEÇOU O SV: \n", err)
	}

}
