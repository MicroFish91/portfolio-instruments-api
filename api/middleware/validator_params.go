package middleware

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func AddParamsValidator[T any]() fiber.Handler {
	return func(c fiber.Ctx) error {
		var params T

		// Initialize a map with all the same fields as the generic params schema T
		pmap := genericToMap(params)

		// Iterate over the map's ids and actually try to fetch those values one-by-one off the request params
		for key := range pmap {
			pmap[key] = c.Params(key)
		}

		// Map the values we found back onto the original params struct so that we can now leverage the schema's Validator
		err := mapAnyToStruct(pmap, &params)
		if err != nil {
			return err
		}

		v, ok := any(params).(Validator)
		if !ok {
			return utils.SendError(c, fiber.StatusInternalServerError, errors.New("internal: request params validation requires an implementation of the Validator interface"))
		}

		if err := v.Validate(); err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, fmt.Errorf("invalid request params: %v", err.Error()))
		}

		c.Locals(constants.LOCALS_REQ_PARAMS, params)
		return c.Next()
	}
}

func genericToMap(g interface{}) map[string]interface{} {
	gobjValue := reflect.ValueOf(g)
	gobjType := gobjValue.Type()

	gmap := map[string]interface{}{}
	for i := 0; i < gobjValue.NumField(); i++ {
		f := gobjValue.Field(i)
		// fmt.Printf("%d: %s %s = %v\n", i, genType.Field(i).Name, f.Type(), f.Interface()) // uncomment this to see what's happening
		gmap[strings.ToLower(gobjType.Field(i).Name)] = f.Interface()
	}

	return gmap
}

func mapAnyToStruct(m map[string]interface{}, target interface{}) error {
	if reflect.ValueOf(target).Kind() != reflect.Ptr {
		return errors.New("internal: the target param struct must be passed in as a pointer")
	}

	for key, value := range m {
		err := setStructField(target, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
