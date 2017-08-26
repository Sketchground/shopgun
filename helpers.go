package shopgun

import "fmt"

// buildPath returns a string in the given format "?a=1&b=2"
func buildPath(base string, args map[string]string) string {
	cnt := 0
	for k, v := range args {
		cnt++
		if cnt == 1 {
			base = fmt.Sprintf("%v?%v=%v", base, k, v)
			continue
		}
		base = fmt.Sprintf("%v&%v=%v", base, k, v)
	}
	return base
}

// stringListArgs builds up a comma separated list of args e.g. a,b,c,d,e
func stringListArgs(args []string) string {
	res := ""
	for i, a := range args {
		if i == 0 {
			res = a
			continue
		}
		res = fmt.Sprintf("%v,%v", res, a)
	}
	return res
}
