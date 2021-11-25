/**
 * @Author: Resynz
 * @Date: 2021/11/25 11:16
 */
package go_ffmpeg_core

type PeelType uint8

const (
	PeelTypeUnknown PeelType = iota
	PeelTypeVideo
	PeelTypeAudio
)
