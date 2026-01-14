package doctor

import (
	"fmt"
)

func RunDoctor() {
	fmt.Println("Dictionary Doctor")
	fmt.Println("-----------------")

	checkOS()
	checkBinary()
	checkConfig()
	checkResources()
	checkNativeMessaging()
	checkLookup()
}
