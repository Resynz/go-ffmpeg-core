/**
 * @Author: Resynz
 * @Date: 2021/11/25 13:57
 */
package go_ffmpeg_core

import "os/exec"

type Transfer struct {
	Command      string
	ResourcePath string
	OutputPath   string
	Quantity     string
	Bitrate      string
}

func (s *Transfer) Execute() error {
	args := make([]string, 0)
	args = append(args, "-i", s.ResourcePath, "-s", s.Quantity, "-c:v", "h264", "-b:v", s.Bitrate, s.OutputPath)
	cmd := exec.Command(s.Command, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		return HandleError(ot)
	}
	return err
}
