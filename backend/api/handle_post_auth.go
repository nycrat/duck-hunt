package api

import (
	"backend/common"
	"database/sql"
	"net/http"
	"strings"
)

func HandlePostAuth(w http.ResponseWriter, r *http.Request) {
	scheme, passcode, found := strings.Cut(r.Header.Get("Authorization"), " ")

	if !found || scheme != "Basic" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	pepper := r.Context().Value("pepper").([]byte)
	db := r.Context().Value("db").(*sql.DB)

	id, ok := common.DbSelectId(passcode, pepper, db)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	hs256Key := r.Context().Value("key").([]byte)

	token := common.GenerateJwtToken(id, hs256Key)
	w.Write([]byte(token))
}
