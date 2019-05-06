package remember

import (
	"fmt"
	"os"
)

func debugLog(data interface{}) {
	if "" != os.Getenv("DEBUG") {
		fmt.Println("DEBUG", data)
	}
}
