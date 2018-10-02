package main

/*
  Import relevant packages.
  HINT: You may need encoding/json package.
*/
import (
  "net/http"
  "exercises/zipserver/handlers"
  "os"
  "fmt"
  "exercises/zipserver/models"
)

func main() {
  mux := http.NewServeMux()

  f, err := os.Open("zips.csv")
  defer f.Close()

  if err != nil {
    fmt.Println("Error couldn't read file.")
    return
  } else {
    zList, err := models.LoadZips(f, 42614)

    if err != nil {
      fmt.Println("Failed to create ZipIndex")
      return
    } else {
      ctx := handlers.Ctx{
        Db: models.BuildIndex(zList),
      }

      // register your handler function
      mux.HandleFunc("/", IndexHandler)
      mux.HandleFunc("/search", ctx.Handler)
      http.ListenAndServe(":4000", mux)
    }
  }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
}
