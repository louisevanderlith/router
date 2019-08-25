package logic

import (
	"testing"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/servicetype"

	uuid "github.com/nu7hatch/gouuid"
)

func dummyService(name string) *bodies.Service {
	return bodies.NewService(name, "", 1, servicetype.API)
	/*  mango.Service{
	Name:        name,
	URL:         "http://127.0.01/" + name,
	Type:        enums.API}*/
}

func TestAddService_ShouldCreateUUID(t *testing.T) {
	service := dummyService("Test.API")

	result, err := AddService(service)

	if err != nil {
		t.Error(err)
	}

	if result == "" {
		t.Error("didn't generate UUID")
	}
}

/*
func TestGetService_AllowedCaller_ForApplication_IsAll(t *testing.T) {
	app := dummyService("Test.APP")
	app.Type = servicetype.APP

	AddService(app)

	result, err := getService("Test.APP", app.Type)

	if err != nil {
		t.Error(err)
		return
	}

	expect := servicetype.ANY

	for _, v := range result.AllowedCallers {
		if v == expect {
			return
		}
	}

	t.Fail("Allowed Called not Found: ")

	if result.AllowedCallers != expect {
		t.Errorf("Allowed Caller is not %s, instead got %s", expect, result.AllowedCallers.String())
	}
}*/

func TestGetServicePath_SameEnv_ShouldFindService(t *testing.T) {
	requestor := dummyService("Test.APP")
	requestor.Type = servicetype.APP
	requestorID, err := AddService(requestor)

	if err != nil {
		t.Error(err)
	}

	api := dummyService("Test.API")
	AddService(api)

	_, err = GetServicePath("Test.API", requestorID, false)

	if err != nil {
		t.Error(err)
	}
}

func TestGetServicePath_DiffEnv_ShouldHaveError(t *testing.T) {
	requestor := dummyService("Test.APP")
	requestor.Type = servicetype.APP
	requestorID, err := AddService(requestor)

	if err != nil {
		t.Error(err)
	}

	api := dummyService("Test.API")
	AddService(api)

	_, err = GetServicePath("Test.API", requestorID, false)

	if err == nil {
		t.Error("Expecting an error message: Test.API wasn't found for the requesting application")
	}
}

func TestGetServicePath_FakeRequestorID_ShouldHaveError(t *testing.T) {
	requestorID, _ := uuid.NewV4()

	api := dummyService("Test.API")
	AddService(api)

	_, err := GetServicePath("Test.API", requestorID.String(), false)

	if err == nil {
		t.Error("Expecting an error message: Couldn't find an application with the given appID")
	}
}

func TestGetServicePath_SameService__CantCallSelf_ShouldHaveError(t *testing.T) {
	requestor := dummyService("Test.API")
	requestor.Type = servicetype.API
	requestorID, err := AddService(requestor)

	if err != nil {
		t.Error(err)
	}

	_, err = GetServicePath("Test.API", requestorID, false)

	if err == nil {
		t.Error("Expecting 'Test.API wasn't found for the requesting application'")
	}
}
