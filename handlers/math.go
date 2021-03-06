package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WilliamMortlMicrosoft/HelloGoService/shared"
	"github.com/WilliamMortlMicrosoft/HelloGoService/telemetry"
	"github.com/gorilla/mux"
)

// mathNumbers is used to load an array of numbers from JSON
type mathNumbers struct {
	Numbers []float64 `json:"numbers"`
}

type operator string

const (
	add      operator = "Add"
	subtract operator = "Subtract"
	multiply operator = "Multiply"
)

// telemetry for math
var telemetryMath *telemetry.Telemetry

// MathHandler is a http handler for math requests
// @Summary mathematics service
// @Description performs 3 operations: Add, Subtract, Multiply
// @Tags advanced services
// @Accept json
// @Produce json
// @Param operator path string true "Add|Subtract|Multiply"
// @Param message body mathNumbers true "numbers"
// @Success 200 {object} mathNumbers
// @Failure 400 "error message"
// @Router /Math/{operator} [post]
func MathHandler() func(w http.ResponseWriter, r *http.Request) {

	// initialize telemetry only on the first call
	if telemetryMath == nil {
		telemetryMath = telemetry.InitializeTelemetryDefault("Math")
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// read JSON body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorTitle := "Missing JSON"
			errorMsg := "JSON was not posted posted to this endpoint!"
			http.Error(w, errorMsg, http.StatusBadRequest)
			(*telemetryMath).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
			return
		}

		// get JSON from the body
		var numbersIn mathNumbers
		err = json.Unmarshal(body, &numbersIn)
		if (err != nil) || (numbersIn.Numbers == nil) || len(numbersIn.Numbers) <= 0 {
			errorTitle := "Malformed JSON"
			errorMsg := "Malformed JSON was posted posted to this endpoint!"
			http.Error(w, errorMsg, http.StatusBadRequest)
			(*telemetryMath).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
			return
		}

		// do the math
		operatorIn := operator(mux.Vars(r)["operator"])
		var operatorFunction func(total float64, number float64) float64
		switch operatorIn {
		case add:
			{
				operatorFunction = func(total float64, number float64) float64 {
					return total + number
				}
				break
			}
		case subtract:
			{
				operatorFunction = func(total float64, number float64) float64 {
					return total - number
				}
				break
			}
		case multiply:
			{
				operatorFunction = func(total float64, number float64) float64 {
					return total * number
				}
				break
			}
		default:
			{
				errorTitle := "Invalid Operator"
				errorMsg := fmt.Sprintf("The operator: %v is not valid!", operatorIn)
				http.Error(w, errorMsg, http.StatusBadRequest)
				(*telemetryDBAdd).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
				return

			}
		}

		// call go routine... this is the equivalent of a go thread
		c := make(chan float64)
		go threadMath(numbersIn.Numbers, operatorFunction, c)

		// wait for channel c to have a response
		var ret mathNumbers
		ret.Numbers = make([]float64, 1)
		ret.Numbers[0] = <-c

		// output the total
		response, _ := json.Marshal(ret)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(response))
		(*telemetryHello).LogInfo("Call", fmt.Sprintf("Math operator: %v IP: %v", operatorIn, shared.GetIP(r)))
	}
}

// threadMath is a go routine that performs math as if it was a separate thread
func threadMath(numbers []float64, operatorFunction func(total float64, number float64) float64, c chan float64) {
	total := numbers[0]
	for i := 1; i < len(numbers); i++ {
		total = operatorFunction(total, numbers[i])
	}
	c <- total
}
