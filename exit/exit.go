package exit

import (
	"fmt"
	"os"
)

func Error(s string) {
	fmt.Printf("【异常退出】%s\n", s)
	os.Exit(1)
}
