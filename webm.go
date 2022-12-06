/**
 * @Author: Resynz
 * @Date: 2022/12/6 09:27
 */
package go_ffmpeg_core

import "os/exec"

/**
  此模块将视频转换为 web视频格式规范 bitrate= 2000kps  h.264
*/

type WebM struct {
	Command      string
	ResourcePath string
	OutputPath   string
}

func (s *WebM) Execute() error {
	args := make([]string, 0)
	args = append(args, "-i", s.ResourcePath, "-c:v", "h264", "-b:v", "2000k", s.OutputPath)
	cmd := exec.Command(s.Command, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		return HandleError(ot)
	}
	return err
}
