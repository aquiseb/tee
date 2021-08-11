// SOURCE : https://github.com/go-bongo/go-dotaccess
// ALTERNATIVE : https://github.com/stretchr/stew/tree/master/objects
// TODO enable to get into slices thanks to the index -> "cities[0].lat"
package tee

import (
	"errors"
	"reflect"
	"strings"

	"github.com/oleiade/reflections"
)

func Set(obj interface{}, prop string, value interface{}) error {
	// Get the array access
	arr := strings.Split(prop, ".")

	// fmt.Println(arr)
	var err error
	var key string
	last, arr := arr[len(arr)-1], arr[:len(arr)-1]
	for _, key = range arr {
		obj, err = GetProperty(obj, key)
		if err != nil {
			return err
		}
	}

	return SetProperty(obj, last, value)
	// return err
}

func SetProperty(obj interface{}, prop string, val interface{}) error {
	if reflect.TypeOf(obj).Kind() == reflect.Map {

		value := reflect.ValueOf(obj)
		value.SetMapIndex(reflect.ValueOf(prop), reflect.ValueOf(val))
		return nil
	}

	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("Object must be a pointer to a struct")
	}
	prop = strings.Title(prop)

	return reflections.SetField(obj, prop, val)
}
