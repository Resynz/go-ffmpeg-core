/**
 * @Author: Resynz
 * @Date: 2021/11/25 11:48
 */
package go_ffmpeg_core

import (
	"fmt"
	"strings"
)

func HandleError(output []byte) error {
	fmt.Println(string(output))
	ss := strings.Split(string(output), "\n")
	if len(ss) == 0 {
		return nil
	}
	return fmt.Errorf(ss[len(ss)-2])
}
