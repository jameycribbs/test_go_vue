package main

import (
	"net/http"

	"github.com/jameycribbs/case_manager_dashboard/queries"
)

func (a *App) getActiveApps(w http.ResponseWriter, r *http.Request) {
	// Placeholder to simulate session value for logged-in user.
	var currentUserID = 194

	RespondWithJSON(w, http.StatusOK, queries.ActiveApps(a.DB, currentUserID))
}
