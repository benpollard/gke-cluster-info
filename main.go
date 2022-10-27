package main

import (
	"encoding/json"
	"net/http"

	"cloud.google.com/go/compute/metadata"
)

type clusterInfo struct {
	Name string
	Zone string
}

func main() {
	inf := new(clusterInfo)
	c := metadata.NewClient(&http.Client{})
	inf.Name, _ = c.InstanceAttributeValue("cluster-name")
	inf.Zone, _ = c.Zone()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsonResp, _ := json.Marshal(inf)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
		return
	})
	http.ListenAndServe(":8080", nil)
}
