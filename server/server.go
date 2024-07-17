package server

import (
	"encoding/json"
	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"net/http"
        "os"
)

type IPNet struct {
	IP string
}

func StartServer() {
	// DB Lookup Here
	db, err := geoip2.Open("data/GeoIP2-City/GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()

        // default to 8080 if not defined
	// Assumption: Port will be presented as in integer and we will prepend :
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }
        port = ":" + port 

	logrus.Infof("starting server on address: %s", port)
	http.HandleFunc("/location/resolve", func(w http.ResponseWriter, r *http.Request) {
		GeoHandler(db, w, r)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}

func GeoHandler(db *geoip2.Reader, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t IPNet
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Println("Answering for", t.IP)
		// ip := net.ParseIP(r.FormValue("ip"))
		ip := net.ParseIP(t.IP)
		record, err := db.City(ip)
		if err != nil {
			_, _ = w.Write([]byte("[]"))
		}
		mm, err := json.Marshal(record)
		_, _ = w.Write(mm)
	default:
		_, _ = w.Write([]byte("[]"))
	}
}
