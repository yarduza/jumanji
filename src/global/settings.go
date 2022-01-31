package global

import (
	"net/http"
)

const LocalhostPort = ":8888"
const ProductionPORT = ":8080"


var CORSOrigin string
var CurrentPort string


func SetDefaultHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", CORSOrigin)
	w.Header().Set("Content-Type", "application/json")
}
