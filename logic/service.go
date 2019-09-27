package logic

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/servicetype"

	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

type Services []*bodies.Service

var serviceMap map[string]Services

func init() {
	serviceMap = make(map[string]Services)
}

func GetServiceMap() map[string]Services {
	return serviceMap
}

func GetApplicants(profile string) map[string]string {
	result := make(map[string]string)

	for _, services := range serviceMap {
		for _, srv := range services {
			log.Println(srv.Profile)
			if srv.Profile == "" || srv.Profile == profile {
				result[srv.Name] = detectSubdomain(srv.Name, profile)
			}
		}
	}

	return result
}

// AddService registers a new service and returns a key for that entry
func AddService(service *bodies.Service) (string, error) {
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

func detectSubdomain(srvcName, profile string) string {
	dotIdx := strings.Index(srvcName, ".")

	return strings.Replace(strings.ToLower(srvcName[:(dotIdx)]), strings.ToLower(profile), "", -1)
}

func isDuplicate(s *bodies.Service) (*bodies.Service, bool) {
	items, _ := serviceMap[s.Name]

	for _, value := range items {
		if value.URL == s.URL {
			return value, true
		}
	}

	return nil, false
}

// GetServicePath will return the correct URL for a requested service.
func GetServicePath(serviceName, appID, host string, clean bool) (string, error) {
	requestingApp := getRequestingService(appID)

	if requestingApp == nil {
		return "", errors.New("Couldn't find an application with the given appID")
	}

	if clean {
		keyName := strings.Split(serviceName, ".")[0]

		return "https://" + strings.ToLower(keyName) + host, nil
	}

	service, err := getService(serviceName, requestingApp.Profile, requestingApp.Type)

	if err != nil {
		return "", fmt.Errorf("%s didn't find %s. %+v", requestingApp.Name, serviceName, err)
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

func getService(serviceName, profile string, callerType servicetype.Enum) (*bodies.Service, error) {
	serviceItems, ok := serviceMap[serviceName]

	if !ok {
		return nil, fmt.Errorf("%s not in serviceMap", serviceName)
	}

	for _, val := range serviceItems {
		profileMatch := val.Profile == "" || val.Profile == profile
		_, allowAny := val.AllowedCallers[servicetype.ANY]

		if allowAny && profileMatch {
			return val, nil
		}

		_, isAllowed := val.AllowedCallers[callerType]

		if isAllowed && profileMatch {
			return val, nil
		}
	}

	return nil, fmt.Errorf("no allowed services available for %v at %s", callerType, profile)
}

func getRequestingService(appID string) *bodies.Service {
	var result *bodies.Service

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
