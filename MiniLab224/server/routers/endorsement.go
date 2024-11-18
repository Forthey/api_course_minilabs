package routers

import (
	"fmt"
	"net/http"
	"server/database/queries"
)

const endorsePrefix = "/endorse"

func endorsement(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		endorseCoin(w, r)
	} else if r.Method == "DELETE" {
		removeEndorsement(w, r)
	}
}

func endorseCoin(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err := queries.Endorse(id)
	if err != nil {
		fmt.Println(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func removeEndorsement(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err := queries.RemoveEndorsement(id)
	if err != nil {
		fmt.Println(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}
