package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cookbook/service"
	"github.com/gorilla/mux"
)

type MealPlanHandler struct {
	Service service.MealPlanService
}

func (handler MealPlanHandler) Register(router *mux.Router) {
	subrouter := router.PathPrefix("/meal-plans").Subrouter()
	subrouter.Path("").Methods(http.MethodGet).HandlerFunc(handler.mealPlans)
	subrouter.Path("").Methods(http.MethodPost).HandlerFunc(handler.createMealPlan)
	subrouter.Path("/{id}").Methods(http.MethodGet).HandlerFunc(handler.mealPlanById)
	subrouter.Path("/{id}").Methods(http.MethodPut).HandlerFunc(handler.updateMealPlan)
	subrouter.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(handler.deleteMealPlan)
}

func (handler MealPlanHandler) mealPlans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mealPlans, err := handler.Service.GetAll()
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(mealPlans)
}

func (handler MealPlanHandler) mealPlanById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	mealPlan, err := handler.Service.Get(id)
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(mealPlan)
}

func (handler MealPlanHandler) createMealPlan(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var mealPlan service.MealPlanCreate
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&mealPlan)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	err = handler.Service.Create(mealPlan)
	if err != nil {
		handleError(w, err)
		return
	}
	errorResponse(w, "Success", http.StatusOK)
}

func (handler MealPlanHandler) updateMealPlan(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var mealPlan service.MealPlanCreate
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&mealPlan)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	err = handler.Service.Update(mealPlan)
	if err != nil {
		handleError(w, err)
		return
	}
	errorResponse(w, "Success", http.StatusOK)
}

func (handler MealPlanHandler) deleteMealPlan(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}
	err = handler.Service.Delete(id)
	if err != nil {
		handleError(w, err)
		return
	}
}
