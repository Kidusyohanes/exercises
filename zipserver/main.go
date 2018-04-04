package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/info344-s18/exercises/zipserver/handlers"
	"github.com/info344-s18/exercises/zipserver/models"
)

func main() {
	//open the zips.csv file and report an errors
	f, err := os.Open("zips.csv")
	if err != nil {
		//if we can't open the file, we can't continue
		//so use log.Fatalf() to report the error and
		//stop the process
		log.Fatalf("error opening zips.csv: %v", err)
	}
	//load the zip code records from the file
	//and report any errors
	zips, err := models.LoadZips(f, 42613)
	if err != nil {
		//again, we can't continue in this case
		//so use log.Fatal()
		log.Fatalf("error parsing CSV: %v", err)
	}
	//close the file so we don't leave it open
	//while the web server is running
	f.Close()
	log.Printf("loaded %d zips", len(zips))

	//build a map index from city name (lower-cased)
	//to a slice of *Zip for that city
	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}
	log.Printf("there are %d zips in Seattle", len(cityIndex["seattle"]))

	//construct a new ZipIndexHandler, passing the
	//cityIndex as the index to use
	cityZipHandler := handlers.NewZipIndexHandler(cityIndex)

	//get the address to listen on from the
	//ADDR environment variable
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")
	if len(tlsKeyPath) == 0 || len(tlsCertPath) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}
	//create a new mux and add two routes:
	// / => root handler
	// /zips/city/<city-name> => cityZipHandler
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.Handle("/zips/city/", cityZipHandler)

	//start the web server
	log.Printf("server is listening at https://%s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}
