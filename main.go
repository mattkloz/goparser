package main

import (
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/mattkloz/goparser/parse"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var (
	serverport = goDotEnvVariable("SERVER_PORT")
	dbip       = goDotEnvVariable("DB_IP_ADDRESS")
	dbport     = goDotEnvVariable("DB_PORT")
	dbname     = goDotEnvVariable("DB_NAME")
	dbuser     = goDotEnvVariable("DB_USER")
	dbpass     = goDotEnvVariable("DB_PASS")
	sqlinit    = dbuser + ":" + dbpass + "@tcp(" + dbip + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _      = gorm.Open("mysql", sqlinit)
)

type ParseItemModel struct {
	Id        int    `gorm:"primary_key"`
	RefUrl    string `gorm:"type:text"`
	PageData  string `gorm:"type:text"`
	Completed bool
}

func AddtoDb(pagedata string, refurl string, w http.ResponseWriter, r *http.Request) string {
	parseditem := &ParseItemModel{PageData: pagedata, RefUrl: refurl, Completed: false}
	db.Create(&parseditem)
	result := db.Last(&parseditem).Value
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return "complete"
}

func CreateItem(w http.ResponseWriter, r *http.Request) {

	// get the values from the post request
	pagedata := r.FormValue("pagedata")
	refurl := r.FormValue("ref_url")

	// send the values to the parser
	getparse := parse.ParseItem(pagedata)

	// if the return is an array, each array item is stored separately. If not an array, the string is stored.
	if len(getparse) > 1 {
		for i := 0; i < len(getparse); i++ {
			AddtoDb(getparse[i], refurl, w, r)
		}
	} else {
		AddtoDb(getparse[0], refurl, w, r)
	}
}

func init() {
	// initialize logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	defer db.Close()

	db.Debug().DropTableIfExists(&ParseItemModel{})
	db.Debug().AutoMigrate(&ParseItemModel{})

	log.Info("Starting GoParser API server")
	router := mux.NewRouter()
	router.HandleFunc("/parse", CreateItem).Methods("POST")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(serverport, handler)
}
