package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/couchbase/gocb"
	"github.com/go-chi/chi"
	"github.com/utschenik/PropertiesApi/db"
)

type languageItem struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}
type property struct {
	Section   string         `json:"section"`
	Languages []languageItem `json:"languages"`
}

type requestProperty struct {
	Section string
	Name    string
}

func (r *requestProperty) DeleteProperty() error {
	_, err := db.Collection.Remove(r.Name, &gocb.RemoveOptions{})

	if err != nil {
		return err
	}
	return nil
}

// GetPropertiesBySection returns all propeties from the bucket for the section
func GetPropertiesBySection(w http.ResponseWriter, r *http.Request) {
	query := fmt.Sprintf("SELECT properties.* FROM properties WHERE section = '%s';", chi.URLParam(r, "section"))
	rows, err := db.Cluster.Query(query, &gocb.QueryOptions{})

	if err != nil {
		panic(err)
	}

	var res []property
	for rows.Next() {
		var row property
		err := rows.Row(&row)
		if err != nil {
			panic(err)
		}
		res = append(res, row)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetPropertyByName returns all properties from a section
func GetPropertyByName(w http.ResponseWriter, r *http.Request) {
	res, err := db.Collection.Get(chi.URLParam(r, "name"), nil)
	if err != nil {
		panic(err)
	}

	var foundProperty property
	err = res.Content(&foundProperty)
	if err != nil {
		panic(err)
	}
	json, err := json.Marshal(foundProperty)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// CreatePropertyForSection creates a new property for a section
func CreatePropertyForSection(w http.ResponseWriter, r *http.Request) {
	newLanguageItem := languageItem{Text: "nt12312323", Language: "spa"}
	newLanguages := []languageItem{newLanguageItem}
	newProperty := property{Section: chi.URLParam(r, "section"), Languages: newLanguages}
	_, err := db.Collection.Insert(chi.URLParam(r, "name"), &newProperty, nil)
	if err != nil {
		panic(err)
	}
}

// DeletePropertyByName delete a property in a section
func DeletePropertyByName(w http.ResponseWriter, r *http.Request) {
	var req = requestProperty{chi.URLParam(r, "section"), chi.URLParam(r, "name")}
	var err = req.DeleteProperty()
	if err != nil {
		panic(err)
	}
}
