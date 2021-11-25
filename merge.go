/**
 * @Author: Resynz
 * @Date: 2021/11/25 11:38
 */
package go_ffmpeg_core

import (
	"os/exec"
)

// MergeAV 合并音视频
type MergeAV struct {
	Command    string
	VideoPath  string
	AudioPath  string
	OutputPath string
}

func (s *MergeAV) Execute() error {
	args := make([]string, 0)
	args = append(args, "-i", s.VideoPath)
	args = append(args, "-i", s.AudioPath)
	args = append(args, "-c", "copy", s.OutputPath)
	cmd := exec.Command(s.Command, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		return HandleError(ot)
	}
	return nil
}
