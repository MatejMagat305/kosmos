package variables

import (
	"fmt"
)

func AddWarning(s string) {
	if len(Warning) > 0 {
		Warning = fmt.Sprint(Warning, "\n", s)
		return
	}
	Warning = s

}
