package devtools

import "strings"

func PathFormat(path string, number int) []string {

	branchArray := strings.SplitN(path, "/", number)
	return branchArray
}
