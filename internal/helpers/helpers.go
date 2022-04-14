package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

// MustMapEnv enables the allocation of an environment variable to a provided target variable in
// order to remove some boilerplate error checking for environment variables that must be set.
func MustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		log.Fatalf("environment variable %q not set", envKey)
	}
	*target = v
}

// SaveAsJSON writes the provided interface to a json file of the given filename.
func SaveAsJSON(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, file, 0644)
}

// GetListOfFields returns a list of fields for a given type, e.g. GetListOfFields(ciscobcs.Device{})
// TODO: Use this later to provide CSV export
func GetListOfFields(v interface{}) []string {
	ret := []string{}
	val := reflect.Indirect(reflect.ValueOf(v))
	numFields := val.Type().NumField()
	for i := 0; i < numFields; i++ {
		ret = append(ret, val.Type().Field(i).Name)
	}
	return ret
}
