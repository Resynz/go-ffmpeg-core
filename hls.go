/**
 * @Author: Resynz
 * @Date: 2022/8/3 17:17
 */
package go_ffmpeg_core

import (
	"fmt"
	"os/exec"
)

type Hls struct {
	Command         string
	InputPath       string
	KeyInfoFile     string
	BaseUrl         string
	HLSTime         uint64
	SegmentFilename string
	OutputPath      string
}

func (s *Hls) Execute() error {
	args := make([]string, 0)
	args = append(args, "-y", "-i", s.InputPath, "-c:v", "libx264", "-c:a", "aac", "-bsf:a", "aac_adtstoasc")
	args = append(args, "-f", "hls", "-hls_playlist_type", "vod")
	if s.KeyInfoFile != "" {
		args = append(args, "-hls_key_info_file", s.KeyInfoFile)
	}
	args = append(args, "-hls_flags", "split_by_time")
	args = append(args, "-hls_time", fmt.Sprintf("%d", s.HLSTime))
	args = append(args, "-hls_segment_filename", s.SegmentFilename)
	if s.BaseUrl != "" {
		args = append(args, "-hls_base_url", s.BaseUrl)
	}
	args = append(args, s.OutputPath)
	cmd := exec.Command(s.Command, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		return HandleError(ot)
	}
	return nil
}
