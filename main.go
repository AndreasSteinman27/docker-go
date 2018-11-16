package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

type Response struct {
    Message string `json:"message"`
}

type TimeZone struct {
    TimeZoneId int
    TimeZoneName string
    DisplayName string
    TzDbName string
    CountryCode int
    Alpha2 string
    CountryName string
}

const(
    GET_TIME_URL = "https://usglobal.cvvplus.com/global.service/test/api/TimeZone/TzDb/us"
    POST_TIME_URL = "https://usglobal.com/test/api/TimeZone/TzDb/us"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
    // make the GET request for timezones
    getResponse, err := http.Get(GET_TIME_URL)

    // handle errors for the GET request
    if err != nil {
        log.Fatal(err)
        http.Error(w, err.Error(), 500)
        return
    }

    // read the body of the GET request
    defer getResponse.Body.Close()
    body, _ := ioutil.ReadAll(getResponse.Body)

    // transform raw bytes into go structs
    // Not necessary, but shows that we could do this.
    var timezones []TimeZone
    json.Unmarshal(body, &timezones)

    // loop through the timezones and do something
    for _, tz := range timezones {
        if (strings.HasPrefix(tz.TimeZoneName, "East")) {
            fmt.Printf("%+v\n", tz)
        }
    }

    // transform structs back into bytes 
    b, err := json.Marshal(timezones)

    if err != nil {
		fmt.Println("error:", err)
	}

    // make a POST request with bytes from marshalled JSON
    postResponse, postError := http.Post(POST_TIME_URL, "application/json", bytes.NewBuffer(b))

    // handler errors from the POST request
    if postError != nil {
        log.Fatal(postError)
        http.Error(w, postError.Error(), 500)
        return
    }

    log.Printf("Received response from  with status: %d", postResponse.StatusCode)

    // Set this header to let the client know what kind of data they expect
    w.Header().Set("Content-Type", "application/json")

    if (postResponse.StatusCode == 200) {
        w.WriteHeader(200)
        r := Response{Message: "Success"}
        json.NewEncoder(w).Encode(r)
    } else {
        // always return 400 if there was an error. 
        w.WriteHeader(400)
        r := Response{Message: "Error"}
        json.NewEncoder(w).Encode(r)
    }
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", GetHandler).Methods("GET")
    log.Printf("Listening on %s...", "8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
