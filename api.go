package testifytutorial

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Location struct {
	Name string
	Lat  float64
	Lng  float64
}

const LocationKind = "Location"

func getLocations(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var locs []Location
	q := datastore.NewQuery(LocationKind)
	_, err := q.GetAll(ctx, &locs)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"Error":"%v"}`, err)))
		return
	}

	res, err := json.Marshal(locs)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"Error":"%v"}`, err)))
		return
	}
	w.Write(res)
}

func addLocation(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	locationKey := datastore.NewIncompleteKey(ctx, LocationKind, nil)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"Error":"%v"}`, err)))
		return
	}

	var loc Location
	err = json.Unmarshal(reqBody, &loc)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"Error":"%v"}`, err)))
		return
	}

	_, err = datastore.Put(ctx, locationKey, &loc)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"Error":"%v"}`, err)))
		return
	}
	w.Write([]byte(`{"addLocation":"success"}`))
}
