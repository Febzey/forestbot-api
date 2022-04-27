package configs

// import (
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/febzey/forestbot-api/pkg/utils"
// 	"github.com/gorilla/mux"
// )

// func ServerConfig(router *mux.Router) *http.Server {
// 	serverConnectionUrl, _ := utils.ConnectionUrlBuilder("server")
// 	readTImeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

// 	return &http.Server{
// 		Handler:     router,
// 		Addr:        serverConnectionUrl,
// 		ReadTimeout: time.Duration(readTImeoutSecondsCount) * time.Second,
// 	}
// }
