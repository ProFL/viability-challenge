package viability

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LocationResponse struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type ResponseSplitter struct {
	ID       string           `json:"id"`
	Location LocationResponse `json:"location"`
}

// Response é o modelo da resposta do método que realiza a viabilidade
type Response struct {
	IsApproved bool             `json:"isApproved"`
	Distance   float64          `json:"distance"`
	Splitter   ResponseSplitter `json:"splitter"`
}

// ClientErrorResponse é o modelo de resposta para requisições de código 4xx
type ClientErrorResponse struct {
	Message string `json:"message"`
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/viability", index).Methods("GET")

	return r
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	latStr := req.URL.Query().Get("lat")
	longStr := req.URL.Query().Get("long")

	if latStr == "" {
		res.WriteHeader(400)
		resStr, err := json.Marshal(ClientErrorResponse{
			"(lat)itude query parameter must be provided",
		})
		if err != nil {
			log.Panic(err)
		}
		res.Write(resStr)
		return
	}

	if longStr == "" {
		res.WriteHeader(400)
		resStr, err := json.Marshal(ClientErrorResponse{
			"(long)itude query parameter must be provided",
		})
		if err != nil {
			log.Panic(err)
		}
		res.Write(resStr)
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)

	if err != nil {
		res.WriteHeader(400)
		resStr, err := json.Marshal(ClientErrorResponse{
			"(lat)itude query parameter must be a float",
		})
		if err != nil {
			log.Panic(err)
		}
		res.Write(resStr)
		return
	}

	long, err := strconv.ParseFloat(longStr, 64)

	if err != nil {
		res.WriteHeader(400)
		resStr, err := json.Marshal(ClientErrorResponse{
			"(long)itude query parameter must be a float",
		})
		if err != nil {
			log.Panic(err)
		}
		res.Write(resStr)
		return
	}

	resStr, err := json.Marshal(Response{
		IsApproved: true,
		Distance:   100,
		Splitter: ResponseSplitter{
			ID: "1234",
			Location: LocationResponse{
				Lat:  lat,
				Long: long,
			},
		},
	})
	if err != nil {
		log.Panic(err)
	}
	res.Write(resStr)
}
