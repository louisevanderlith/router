package logic

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

type Services []*mango.Service

var serviceMap map[string]Services

func init() {
	serviceMap = make(map[string]Services)
}

func GetServiceMap() map[string]Services {
	return serviceMap
}

// AddService registers a new service and returns a key for that entry
func AddService(service *mango.Service) (string, error) {
	if !strings.Contains(service.Name, ".") {
		return "", errors.New("invalid service Name")
	}

	val, duplicate := isDuplicate(service)

	if duplicate {
		val.Version++
		return val.ID, nil
	}

	u4, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	service.ID = u4.String()
	service.Version = getVersion()
	service.AllowedCaller = getAllowedCaller(service.Type)

	serviceMap[service.Name] = append(serviceMap[service.Name], service)

	return service.ID, nil
}

func isDuplicate(s *mango.Service) (*mango.Service, bool) {
	items, _ := serviceMap[s.Name]

	for _, value := range items {
		if value.URL == s.URL && value.Environment == s.Environment {
			return value, true
		}
	}

	return nil, false
}

// GetServicePath will return the correct URL for a requested service.
func GetServicePath(serviceName, appID string, clean bool) (string, error) {
	requestingApp := getRequestingService(appID)

	if requestingApp == nil {
		return "", errors.New("Couldn't find an application with the given appID")
	}

	if clean {
		keyName := strings.Split(serviceName, ".")[0]
		cleanHost := getCleanHost(requestingApp.Environment)

		return "https://" + strings.ToLower(keyName) + cleanHost, nil
	}

	service, err := getService(serviceName, requestingApp.Environment, requestingApp.Type)

	if err != nil {
		return "", fmt.Errorf("not found. %+v", err)
	}

	return service.URL, nil
}

func getCleanHost(env enums.Environment) string {
	envHost := fmt.Sprintf("HOST_%s", env)

	if len(envHost) == 0 {
		envHost = "HOST_DEV"
	}

	return beego.AppConfig.String(envHost)
}

func getAllowedCaller(serviceType enums.ServiceType) enums.ServiceType {
	var result enums.ServiceType

	switch serviceType {
	case enums.API:
		result = enums.APP
	case enums.APP:
		result = enums.ANY
	}

	return result
}

func getService(serviceName string, environment enums.Environment, callerType enums.ServiceType) (*mango.Service, error) {
	var result *mango.Service
	serviceItems, ok := serviceMap[serviceName]

	if !ok {
		return nil, fmt.Errorf("%s not in serviceMap", serviceName)
	}

	for _, val := range serviceItems {
		correctEnv := val.Environment == environment
		isAllowed := val.AllowedCaller == enums.ANY || val.AllowedCaller == callerType

		if correctEnv && isAllowed {
			result = val
			break
		}
	}

	return result, nil
}

func getRequestingService(appID string) *mango.Service {
	var result *mango.Service

	for _, serviceItems := range serviceMap {
		for _, val := range serviceItems {
			if val.ID == appID {
				result = val
				break
			}
		}
	}

	return result
}

func getVersion() int {
	now := time.Now()
	concatDate := now.Format("0612")
	result, _ := strconv.Atoi(concatDate)

	return result
}
