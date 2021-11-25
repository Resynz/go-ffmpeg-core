/**
 * @Author: Resynz
 * @Date: 2021/11/25 10:40
 */
package go_ffmpeg_core

import "os/exec"

// Peel 剥离音视频
type Peel struct {
	ResourcePath string // 资源路径
	Command      string
	OutputPath   string
	Type         PeelType
	Stream       string
}

func (s *Peel) Execute() error {
	args := make([]string, 0)
	args = append(args, "-i")
	args = append(args, s.ResourcePath)

	if s.Type == PeelTypeVideo {
		args = append(args, "-an", "-vcodec")
	}
	if s.Type == PeelTypeAudio {
		args = append(args, "-vn", "-acodec")
		if s.Stream != "" {
			args = append(args, "-map", s.Stream)
		}
	}

	args = append(args, "copy", s.OutputPath)
	cmd := exec.Command(s.Command, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		return HandleError(ot)
	}
	return nil
}
