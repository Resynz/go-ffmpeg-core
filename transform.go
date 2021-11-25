/**
 * @Author: Resynz
 * @Date: 2021/11/25 13:57
 */
package go_ffmpeg_core

import "os/exec"

type Transform struct {
	Command      string
	ResourcePath string
	OutputPath   string
}

func (s *Transform) Execute() error {
	args := make([]string, 0)
	args = append(args, "-i", s.ResourcePath, s.OutputPath)
	cmd := exec.Command(s.Command, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		return HandleError(ot)
	}
	return err
}
