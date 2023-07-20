package helpers

import "fmt"

type Option string

const (
	Panic Option = "panic"
	Log   Option = "log"
)

func ValidateOption(opt Option) bool {
	return opt == Panic || opt == Log
}

func HandleError(err error, opt Option) (error error) {
	if ValidateOption(opt) {
		if opt == Panic {
			panic(err)
			return err
		} else {
			fmt.Println(err)
			return err
		}
	} else {
		panic("Invalid option passed to HandleError function")
		return err
	}
}
