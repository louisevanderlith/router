package logic

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/servicetype"

	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

type Services []*droxolite.Service

var serviceMap map[string]Services

func init() {
	serviceMap = make(map[string]Services)
}

func GetServiceMap() map[string]Services {
	return serviceMap
}

// AddService registers a new service and returns a key for that entry
func AddService(service *droxolite.Service) (string, error) {
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
	service.AllowedCallers = getAllowedCaller(service.Type)

	serviceMap[service.Name] = append(serviceMap[service.Name], service)

	return service.ID, nil
}

func isDuplicate(s *droxolite.Service) (*droxolite.Service, bool) {
	items, _ := serviceMap[s.Name]

	for _, value := range items {
		if value.URL == s.URL {
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
		cleanHost := os.Getenv("HOST")

		return "https://" + strings.ToLower(keyName) + cleanHost, nil
	}

	service, err := getService(serviceName, requestingApp.Type)

	if err != nil {
		return "", fmt.Errorf("%s not found. %+v", serviceName, err)
	}

	return service.URL, nil
}

func getAllowedCaller(serviceType servicetype.Enum) map[servicetype.Enum]struct{} {
	result := make(map[servicetype.Enum]struct{})
	if serviceType == servicetype.APP {
		result[servicetype.ANY] = struct{}{}
		return result
	}

	if serviceType == servicetype.APX {
		result[servicetype.APP] = struct{}{}
		return result
	}

	if serviceType == servicetype.API {
		result[servicetype.APX] = struct{}{}
		result[servicetype.APP] = struct{}{}
		return result
	}

	return result
}

func getService(serviceName string, callerType servicetype.Enum) (*droxolite.Service, error) {
	serviceItems, ok := serviceMap[serviceName]

	if !ok {
		return nil, fmt.Errorf("%s not in serviceMap", serviceName)
	}

	for _, val := range serviceItems {
		_, allowAny := val.AllowedCallers[servicetype.ANY]

		if allowAny {
			return val, nil
		}

		_, isAllowed := val.AllowedCallers[callerType]

		if isAllowed {
			return val, nil
		}
	}

	return nil, errors.New("no allowed services available")
}

func getRequestingService(appID string) *droxolite.Service {
	var result *droxolite.Service

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
