package testifytutorial

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestAddLocation(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatalf("Error creating aetest instance: %v", err)
	}
	defer inst.Close() // Make sure the aetest instance always closes
	rt := initRouter() // Get our router

	loc, err := json.Marshal(&Location{
		Name: "Cambridge Fresh Pond",
		Lat:  42.385658,
		Lng:  -71.149308,
	})
	if err != nil {
		t.Errorf("Error marshalling Location into JSON: %v", err)
	}

	req, err := inst.NewRequest("POST", "/add-location",
		ioutil.NopCloser(bytes.NewBuffer(loc)))
	if err != nil {
		t.Fatalf("Error preparing request: %v", err)
	}
	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, req)

	const expectedResponse = `{"addLocation":"success"}`
	if string(rec.Body.Bytes()) != expectedResponse {
		t.Errorf("Expected response to be %s, got %s",
			expectedResponse, string(rec.Body.Bytes()))
	}

	dbReq, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Error preparing request: %v", err)
	}
	ctx := appengine.NewContext(dbReq)

	q := datastore.NewQuery("Location")
	numLocs, err := q.Count(ctx)
	if err != nil {
		t.Fatalf("Error preparing request: %v", err)
	}
	if numLocs != 1 {
		t.Errorf("Expected number of locations to be 1, got %d", numLocs)
	}
}

func TestAddLocationTestify(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	require.Nil(t, err, "Error creating aetest instance: %v", err)
	defer inst.Close()
	rt := initRouter()

	loc, err := json.Marshal(&Location{
		Name: "Cambridge Fresh Pond",
		Lat:  42.5,
		Lng:  -71.5,
	})
	assert.Nil(t, err, "Error marshalling Location into JSON: %v", err)

	req, err := inst.NewRequest("POST", "/add-location",
		ioutil.NopCloser(bytes.NewBuffer(loc)))
	require.Nil(t, err, "Error preparing request: %v", err)
	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, req)

	const expectedResponse = `{"addLocation":"success"}`
	assert.Equal(t, expectedResponse, string(rec.Body.Bytes()),
		"Expected response to be %s, got %s", expectedResponse, string(rec.Body.Bytes()))

	dbReq, err := inst.NewRequest("GET", "/", nil)
	require.Nil(t, err, "Error preparing request: %v", err)
	ctx := appengine.NewContext(dbReq)

	q := datastore.NewQuery("Location")
	numLocs, err := q.Count(ctx)
	require.Nil(t, err, "Error preparing request: %v", err)
	assert.Equal(t, 1, numLocs, "Expected number of locations to be 1, got %d", numLocs)
}

func assertExpectedGot(t *testing.T, expected, got interface{}, msg string) {
	assert.Equal(t, expected, got, msg, expected, got)
}

func TestAddLocationExpectedGot(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	require.Nil(t, err, "Error creating aetest instance: %v", err)
	defer inst.Close()
	rt := initRouter()

	loc, err := json.Marshal(&Location{
		Name: "Cambridge Fresh Pond",
		Lat:  42.5,
		Lng:  -71.5,
	})
	assert.Nil(t, err, "Error marshalling Location into JSON: %v", err)

	req, err := inst.NewRequest("POST", "/add-location",
		ioutil.NopCloser(bytes.NewBuffer(loc)))
	require.Nil(t, err, "Error preparing request: %v", err)
	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, req)

	const expectedResponse = `{"addLocation":"success"}`
	assertExpectedGot(t, expectedResponse, string(rec.Body.Bytes()),
		"Expected response to be %s, got %s")

	dbReq, err := inst.NewRequest("GET", "/", nil)
	require.Nil(t, err, "Error preparing request: %v", err)
	ctx := appengine.NewContext(dbReq)

	q := datastore.NewQuery("Location")
	numLocs, err := q.Count(ctx)
	require.Nil(t, err, "Error preparing request: %v", err)
	assertExpectedGot(t, 1, numLocs, "Expected number of locations to be %d, got %d")
}

type TestSuite struct {
	suite.Suite
	inst aetest.Instance
	rt   *mux.Router
}

func (s *TestSuite) SetupSuite() {
	var err error
	s.inst, err = aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	require.Nil(s.T(), err, "Error creating aetest instance: %v", err)
	s.rt = initRouter()
}

func (s *TestSuite) TearDownSuite() {
	err := s.inst.Close()
	assert.Nil(s.T(), err, "Error closing aetest instance: %v", err)
}

func (s *TestSuite) TestAddLocation() {
	t := s.T()

	loc, err := json.Marshal(&Location{
		Name: "Cambridge Fresh Pond",
		Lat:  42.5,
		Lng:  -71.5,
	})
	assert.Nil(t, err, "Error marshalling Location into JSON: %v", err)

	req, err := s.inst.NewRequest("POST", "/add-location",
		ioutil.NopCloser(bytes.NewBuffer(loc)))
	require.Nil(t, err, "Error preparing request: %v", err)
	rec := httptest.NewRecorder()
	s.rt.ServeHTTP(rec, req)

	const expectedResponse = `{"addLocation":"success"}`
	assert.Equal(t, expectedResponse, string(rec.Body.Bytes()),
		"Expected response to be %s, got %s", expectedResponse, string(rec.Body.Bytes()))

	dbReq, err := s.inst.NewRequest("GET", "/", nil)
	require.Nil(t, err, "Error preparing request: %v", err)
	ctx := appengine.NewContext(dbReq)

	q := datastore.NewQuery("Location")
	numLocs, err := q.Count(ctx)
	require.Nil(t, err, "Error preparing request: %v", err)
	assert.Equal(t, 1, numLocs, "Expected number of locations to be 1, got %d", numLocs)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
