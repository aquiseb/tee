// SOURCE : https://github.com/go-bongo/go-dotaccess
// ALTERNATIVE : https://github.com/stretchr/stew/tree/master/objects
// TODO enable to get into slices thanks to the index -> "cities[0].lat"
package tee

import (
	"reflect"
	"strings"

	"github.com/oleiade/reflections"
)

func Get(obj interface{}, prop string) (interface{}, error) {
	// fmt.Println("getting property")
	// fmt.Println(args)

	// Get the array access
	arr := strings.Split(prop, ".")

	// fmt.Println(arr)
	var err error
	// last, arr := arr[len(arr)-1], arr[:len(arr)-1]
	for _, key := range arr {
		obj, err = GetProperty(obj, key)
		if err != nil {
			return nil, err
		}
		if obj == nil {
			return nil, nil
		}
	}
	return obj, nil
}

// Loop through this to get properties via dot notation
func GetProperty(obj interface{}, prop string) (interface{}, error) {

	if reflect.TypeOf(obj).Kind() == reflect.Map {

		val := reflect.ValueOf(obj)

		valueOf := val.MapIndex(reflect.ValueOf(prop))

		if valueOf == reflect.Zero(reflect.ValueOf(prop).Type()) {
			return nil, nil
		}

		idx := val.MapIndex(reflect.ValueOf(prop))

		if !idx.IsValid() {
			return nil, nil
		}
		return idx.Interface(), nil
	}

	prop = strings.Title(prop)
	return reflections.GetField(obj, prop)
}
