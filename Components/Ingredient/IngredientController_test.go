package Ingredient

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Test"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ingredientController *IngredientController

func init() {
	Test.OpenTestingDatabase()
	ingredientController = NewIngredientController(NewIngredientService(NewIngredientRepository()))
}

func TestIngredientController_GetAllIngredient(t *testing.T) {

	router := gin.Default()

	router.GET("/api/v1/ingredient", ingredientController.GetAllIngredient)

	req, _ := http.NewRequest("GET", "/api/v1/ingredient", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		fmt.Println(statusOK)

		return statusOK
	})
}

func TestNewIngredientController(t *testing.T) {

	router := gin.Default()

	router.POST("/api/v1/ingredient", ingredientController.CreateIngredient)

	req, _ := http.NewRequest("POST", "/api/v1/ingredient", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		fmt.Println(statusOK)

		return statusOK
	})
}

func testHTTPResponse(t *testing.T, router *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	router.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}
