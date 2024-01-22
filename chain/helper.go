package chain 

import (
	"fmt"
)

func formatErrorMessage(err error) string {
    if err == nil {
        return "No error"
    }
    return fmt.Sprintf("%v", err.Error())
}