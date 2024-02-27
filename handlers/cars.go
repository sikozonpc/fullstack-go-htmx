package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/fullstackgo/components"
	"github.com/sikozonpc/fullstackgo/types"
	"github.com/sikozonpc/fullstackgo/views"
)

func (h *Handler) HandleListCars(w http.ResponseWriter, r *http.Request) {
	isAddingCar := r.URL.Query().Get("isAddingCar") == "true"

	cars, err := h.store.GetCars()
	if err != nil {
		log.Println(err)
		return
	}

	views.Cars(cars, isAddingCar).Render(r.Context(), w)
}

func (h *Handler) HandleAddCar(w http.ResponseWriter, r *http.Request) {
	car := &types.Car{
		Brand:    r.FormValue("brand"),
		Make:     r.FormValue("model"),
		Model:    r.FormValue("make"),
		Year:     r.FormValue("year"),
		ImageURL: r.FormValue("imageURL"),
	}

	newCar, err := h.store.CreateCar(car)
	if err != nil {
		log.Println(err)
		return
	}

	components.CarTile(newCar).Render(r.Context(), w)
}

func (h *Handler) HandleSearchCar(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("search")

	cars, err := h.store.FindCarsByNameMakeOrBrand(text)
	if err != nil {
		log.Println(err)
		return
	}

	components.CarsList(cars).Render(r.Context(), w)
}

func (h *Handler) HandleDeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.store.DeleteCar(id)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
