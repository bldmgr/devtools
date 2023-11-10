package devtools

import "strings"

func pathFormat(path string, number int) []string {

	branchArray := strings.SplitN(path, "/", number)
	return branchArray
}
