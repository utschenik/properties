package handlers

import (
	"encoding/json"
	"net/http"

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

// GetAllProperties returns all propeties from the bucket
func GetAllProperties(w http.ResponseWriter, r *http.Request) {
}

// GetPropertiesBySectionAndName returns all properties from a section
func GetPropertiesBySectionAndName(w http.ResponseWriter, r *http.Request) {
	res, err := db.Collection.Get(chi.URLParam(r, "name"), nil)
	if err != nil {
		panic(err)
	}

	var myProperty property
	err = res.Content(&myProperty)
	if err != nil {
		panic(err)
	}
	json, err := json.Marshal(myProperty)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// CreatePropertyForSection creates a new property for a section
func CreatePropertyForSection(w http.ResponseWriter, r *http.Request) {
	newLanguageItem := languageItem{Text: "nt", Language: "ger"}
	newLanguages := []languageItem{newLanguageItem}
	newProperty := property{Section: chi.URLParam(r, "section"), Languages: newLanguages}
	_, err := db.Collection.Insert(chi.URLParam(r, "name"), &newProperty, nil)
	if err != nil {
		panic(err)
	}
}

// GetPropertiesBySectionAndLanguage returns properties specified on language and section
func GetPropertiesBySectionAndLanguage(w http.ResponseWriter, r *http.Request) {
	var param string = chi.URLParam(r, "language")
	var param1 string = chi.URLParam(r, "section")
}
