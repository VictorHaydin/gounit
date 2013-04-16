package assert

import (
	"fmt"
)

type TeamCity struct {
}

func (t *TeamCity) reportError(file string, line int, message string) {
	fmt.Printf("##teamcity[testFailed name='%v:%v' message='%v']\n", file, line, message)
}

func (t *TeamCity) reportSuccess(file string, line int) {
	fmt.Printf("##teamcity[testFinished name='%v:%v']\n", file, line)
}