package function_2

import (
	"fmt"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/takeoff-pharatesid/cloud-function-poc/common/logging"
	"net/http"
)

func init() {
	logging.NewLogger("function_2")
	functions.HTTP("function2", function2)
}

func function2(w http.ResponseWriter, r *http.Request) {
	logger := logging.NewLogger("function_2")
	logger.Info("Function 2 initialised successfully")
	logger.Debugln("ambassadressAssadAsd")
	logger.Error("This is a n rooerer")
	fmt.Fprintf(w, "Hello, From Function_2 !")
}
