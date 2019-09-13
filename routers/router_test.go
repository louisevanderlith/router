package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/louisevanderlith/droxolite/element"

	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/router/logic"
)

var apiEpoxy resins.Epoxi

func init() {
	srvc := bodies.NewService("Router.API", "/certs/none.pem", 8080, servicetype.API)
	srvc.ID = "RouterTester"

	apiEpoxy = resins.NewMonoEpoxy(srvc, element.GetNoTheme(".localhost/", srvc.ID, "nobody"))
	Setup(apiEpoxy)
}

func TestDiscovery_POST_OK(t *testing.T) {
	servc := bodies.NewService("Nothing.API", "/certs/none.pem", 8095, servicetype.API)
	obj, err := json.Marshal(servc)
	req, err := http.NewRequest(http.MethodPost, "/discovery", bytes.NewBuffer(obj))

	if err != nil {
		t.Fatal(err)
	}

	handle := apiEpoxy.Router()

	rr := httptest.NewRecorder()
	handle.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatal(rr.Code, rr.Body.String())
	}

	result := ""
	_, err = bodies.MarshalToResult(rr.Body.Bytes(), &result)

	if err != nil {
		t.Fatal(err)
	}

	//t.Error(result)

	if len(result) == 0 {
		t.Fatal("result was empty")
	}

}

func TestDiscovery_GET_CleanAppServiceName_InvalidAPP(t *testing.T) {
	///:appID/:serviceName/:clean
	req, err := http.NewRequest("GET", "/discovery/9cfa3c82-88bb-416e-6ba2-34f611ad9e03/Router.API/true", nil)

	if err != nil {
		t.Fatal(err)
	}

	handle := apiEpoxy.Router()

	rr := httptest.NewRecorder()
	handle.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatal(rr.Body.String())
	}

	result := ""
	rest, err := bodies.MarshalToResult(rr.Body.Bytes(), &result)

	if err != nil {
		t.Fatal(err)
	}

	if rest.Reason != "Couldn't find an application with the given appID" {
		t.Fatal("result was empty")
	}
}

func TestDiscovery_GET_DirtyAppServiceName_InvalidAPP(t *testing.T) {
	///:appID/:serviceName/
	req, err := http.NewRequest("GET", "/discovery/9cfa3c82-88bb-416e-6ba2-34f611ad9e03/Router.API", nil)

	if err != nil {
		t.Fatal(err)
	}

	handle := apiEpoxy.Router()

	rr := httptest.NewRecorder()
	handle.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatal(rr.Body.String())
	}

	result := ""
	rest, err := bodies.MarshalToResult(rr.Body.Bytes(), &result)

	if err != nil {
		t.Fatal(err)
	}

	if rest.Reason != "Couldn't find an application with the given appID" {
		t.Fatalf("Not expected: %s", result)
	}
}

func TestMemory_GET_UnAuthed(t *testing.T) {
	req, err := http.NewRequest("GET", "/memory", nil)

	if err != nil {
		t.Fatal(err)
	}

	handle := apiEpoxy.Router()

	rr := httptest.NewRecorder()
	handle.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatal(rr.Body.String())
	}

	result := make(map[string]logic.Services)
	_, err = bodies.MarshalToResult(rr.Body.Bytes(), &result)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMemory_GET_Apps_AuthFails(t *testing.T) {
	req, err := http.NewRequest("GET", "/memory/apps", nil)

	if err != nil {
		t.Fatal(err)
	}

	handle := apiEpoxy.Router()

	rr := httptest.NewRecorder()
	handle.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatal(rr.Body.String())
	}

	result := make(map[string]logic.Services)
	_, err = bodies.MarshalToResult(rr.Body.Bytes(), &result)

	if err != nil {
		t.Fatal(err)
	}
}
