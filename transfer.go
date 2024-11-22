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
	// 确保输出路径以 .mp4 结尾
	if s.OutputPath[len(s.OutputPath)-4:] != ".mp4" {
		s.OutputPath += ".mp4"
	}

	args := []string{
		"-i", s.ResourcePath,
		"-s", s.Quantity,
		"-c:v", "libx264", // H.264 编解码器
		"-b:v", s.Bitrate, // 视频比特率
		"-c:a", "aac", // 音频编解码器 (AAC)
		"-strict", "experimental", // 对于某些 FFmpeg 版本需要添加该参数以启用 AAC
		s.OutputPath,
	}

	cmd := exec.Command(s.Command, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		return HandleError(ot)
	}
	return nil
}
