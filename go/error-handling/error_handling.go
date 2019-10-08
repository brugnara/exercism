package erratum

import (
	"errors"
	"fmt"
	"reflect"
)

// Use will try to open a resource using the ResourceOpener. Depending on
// encountered errors, we will act accordingly. See README.md for better
// understand the exercise.
func Use(f ResourceOpener, str string) (err error) {
	fmt.Println("Use()")
	resource, err := f()

	if err != nil {
		// keep trying if error is type TransientError since the resource
		// will be available soon..
		for fmt.Sprintf("%s", reflect.TypeOf(err)) == "erratum.TransientError" {
			fmt.Println("retrying due erratum.TransientError")
			resource, err = f()
		}
	}

	if err != nil {
		fmt.Printf("Errored with: %v\n", err)
		return
	}

	// if opened, then ensure to always Close the resource
	defer resource.Close()

	// Frob may panic. If the panic contains an error of type FrobError,
	// call the Unfrob(err.defrobTag). Use should then return the error
	// In order to be able to return the error within the deferred func, we
	// must use named parameter in the main func, see: (err error).
	// Changing err value will reflect in the wanted response to the main func.
	defer func(resource Resource) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			errType := fmt.Sprintf("%v", reflect.TypeOf(r))
			errStr := fmt.Sprintf("%v", r)
			fmt.Println(errType)
			fmt.Println(errStr)
			if errType == "erratum.FrobError" {
				frobError, ok := r.(FrobError)
				if !ok {
					fmt.Println("ops. Never happens in this exercise.")
				}
				//var frobError FrobError{r}
				fmt.Println("Defrobbing")
				resource.Defrob(frobError.defrobTag)
			}
			fmt.Printf("Erroring with: %v\n", errStr)
			err = errors.New(errStr)
		}
	}(resource)

	fmt.Println("Before Frob")
	resource.Frob(str)
	fmt.Println("After Frob")

	// no errors
	return
}
