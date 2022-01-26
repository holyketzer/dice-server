package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
)

func random(w http.ResponseWriter, req *http.Request) {
    mins, ok1 := req.URL.Query()["min"]
    maxs, ok2 := req.URL.Query()["max"]

    if ok1 && ok2 {
        min, err1 := strconv.Atoi(mins[0])
        max, err2 := strconv.Atoi(maxs[0])

        if err1 == nil && err2 == nil {
            value := rand.Intn(max - min) + min
            fmt.Println("random number =", value)
            w.Header().Set("Content-Type", "text/html; charset=utf-8")
            response := fmt.Sprintf("<html><body>%d</body></html>", value)
            fmt.Fprint(w, response)
            return
        }
    }

    w.WriteHeader(http.StatusBadRequest)
    w.Header().Set("Content-Type", "application/html")
}

func main() {
    http.HandleFunc("/randint", random)
    http.ListenAndServe(":8888", nil)
}
