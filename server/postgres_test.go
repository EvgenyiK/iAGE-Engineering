package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker pool: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "latest", []string{"POSTGRES_DB=test", "POSTGRES_USER=test", "POSTGRES_PASSWORD=pass"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		dsnString := "postgres://test:pass@0.0.0.0:" + resource.GetPort("5432/tcp") + "/test?sslmode=disable"
		db, err := sqlx.Connect("postgres", dsnString)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Printf("Could not connect to docker resource: %s", err)
	}

	fmt.Println(" Image port: " + resource.GetPort("5432/tcp"))

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		reqBody  []byte
		respBody []byte
	}{
		{[]byte(`{"Name":"Hen","Age":12}`), []byte(`could not create user`)},
		{[]byte(`{"Name":"Maggie","Age":10}`), []byte(`{"ID":12,"Name":"Maggie","Age":10}`)},
		{[]byte(`{"Name":"Maggie","Age":"10"}`), []byte(`invalid body`)},
	}
	for i, v := range testcases {
		req := httptest.NewRequest("POST", "/postgres/users", bytes.NewReader(v.reqBody))
		w := httptest.NewRecorder()

		CreateUser(w,req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.respBody)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.respBody))
		}
	}
}