package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) getLinks(writer http.ResponseWriter, request *http.Request) {
	links, err := app.db.ListAllLinks(app.ctx)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}

	value, err := json.MarshalIndent(map[string]any{"links": links}, "", "  ")
	if err != nil {
		writer.Write([]byte("An error while parsing JSON body"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(value)
}

func (app *application) getLink(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}

	link, err := app.db.GetLink(app.ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Fprintf(writer, "Error: No data found...")
		}

		return
	}

	value, err := json.MarshalIndent(map[string]any{"link": link}, "", "  ")
	if err != nil {
		writer.Write([]byte("An error while parsing JSON body"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(value)
}

func (app *application) deleteLink(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}

	if err := app.db.DeleteLink(app.ctx, int32(id)); err != nil {
		writer.Write([]byte("Could not delete a link with that id"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}
