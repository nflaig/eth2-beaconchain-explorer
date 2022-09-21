package handlers

import (
	"eth2-exporter/templates"
	"net/http"
)

// Faq will return the data from the frequently asked questions (FAQ) using a go template
func UnitConverter(w http.ResponseWriter, r *http.Request) {
	var unitConverterTemplate = templates.GetTemplate("layout.html", "unitConverter.html")

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "unitConverter", "/unitConerter", "Unit Converter")

	err := unitConverterTemplate.ExecuteTemplate(w, "layout", data)

	if err != nil {
		logger.Errorf("error executing template for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusServiceUnavailable)
		return
	}
}
