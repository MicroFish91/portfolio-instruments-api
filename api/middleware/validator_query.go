package middleware

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func AddQueryValidator[T any]() fiber.Handler {
	return func(c fiber.Ctx) error {
		qmap := c.Queries()

		var query T
		if err := mapStringToStruct(qmap, &query); err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, errors.New("error parsing query params"))
		}

		v, ok := any(query).(Validator)
		if !ok {
			return utils.SendError(c, fiber.StatusInternalServerError, errors.New("internal: query params validation requires an implementation of the Validator interface"))
		}

		if err := v.Validate(); err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, fmt.Errorf("invalid query params: %v", err.Error()))
		}

		c.Locals(constants.LOCALS_REQ_QUERY, query)
		return c.Next()
	}
}

func mapStringToStruct(m map[string]string, target interface{}) error {
	if reflect.ValueOf(target).Kind() != reflect.Ptr {
		return errors.New("internal: the target query struct must be passed in as a pointer")
	}

	for key, value := range m {
		err := setStructField(target, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
