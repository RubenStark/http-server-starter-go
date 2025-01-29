package server

import (
	"fmt"
	"net/http"

	"github.com/codecrafters-io/http-server-starter-go/app/routes"
)

func StartServer(
	baseDir string,
) {
	router := routes.NewRouter(baseDir)
	printASCIIArt()
	fmt.Println("Server started running on http://localhost:4221")
	http.ListenAndServe(":4221", router)
}

func printASCIIArt() {
	fmt.Println(`

  ________       .__                                                                  
 /  _____/  ____ |  | _____    ____    ____     ______ ______________  __ ___________ 
/   \  ___ /  _ \|  | \__  \  /    \  / ___\   /  ___// __ \_  __ \  \/ // __ \_  __ \
\    \_\  (  <_> )  |__/ __ \|   |  \/ /_/  >  \___ \\  ___/|  | \/\   /\  ___/|  | \/
 \______  /\____/|____(____  /___|  /\___  /  /____  >\___  >__|    \_/  \___  >__|   
        \/                 \/     \//_____/        \/     \/                 \/       

`)
}
