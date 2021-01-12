package register

import (
	getrandomnumber "github.com/tfpmotta/sre-1/cmd/get-random-number"
	"github.com/tfpmotta/sre-1/internal/router"
)

func init() {
	router.Router.HandleFunc("/get-random-number", getrandomnumber.GetRandomNumber)
}
