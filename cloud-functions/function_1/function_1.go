package function_1

import (
	"fmt"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/takeoff-pharatesid/cloud-function-poc/common/helper"
	"github.com/takeoff-pharatesid/cloud-function-poc/common/logging"
	"net/http"
	"strconv"
)

func init() {
	logging.NewLogger("function_1")
	functions.HTTP("Function1", function1)
}

func sum(a, b int) int {
	return helper.Sum(a, b)
}

func function1(w http.ResponseWriter, r *http.Request) {
	logger := logging.NewLogger("function_1")
	logger.Info("Function 1 initialised successfully")
	logger.Debug("Printing Sum = " + strconv.Itoa(helper.Sum(3, 2)))
	fmt.Fprintf(w, "Hello, From Function_1 !")
}
