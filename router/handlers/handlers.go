package handlers

import "net/http"

// GetAllProperties returns all propeties from the bucket
func GetAllProperties(w http.ResponseWriter, r *http.Request) {}

// GetPropertiesBySection returns all properties from a section
func GetPropertiesBySection(w http.ResponseWriter, r *http.Request) {}

// CreatePropertyForSection creates a new property for a section
func CreatePropertyForSection(w http.ResponseWriter, r *http.Request) {}

// GetPropertiesBySectionAndLanguage returns properties specified on language and section
func GetPropertiesBySectionAndLanguage(w http.ResponseWriter, r *http.Request) {}
