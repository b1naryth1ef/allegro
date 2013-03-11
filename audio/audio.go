package audio

// #cgo LDFLAGS: -lallegro_audio
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_audio.h>
import "C"

import (
//	"github.com/tapir/allegro"
	"unsafe"
)

const (
	AudioDepthInt8 = C.ALLEGRO_AUDIO_DEPTH_INT8
	AudioDepthInt16 = C.ALLEGRO_AUDIO_DEPTH_INT16
	AudioDepthInt24 = C.ALLEGRO_AUDIO_DEPTH_INT24
	AudioDepthFloat32 = C.ALLEGRO_AUDIO_DEPTH_FLOAT32
	AudioDepthUnsigned = C.ALLEGRO_AUDIO_DEPTH_UNSIGNED
	AudioDepthUint8 = C.ALLEGRO_AUDIO_DEPTH_UINT8
	AudioDepthUint16 = C.ALLEGRO_AUDIO_DEPTH_UINT16
	AudioDepthUint24 = C.ALLEGRO_AUDIO_DEPTH_UINT24
)

const AudioPanNone float32 = 1000

const (
	ChannelConf1 = C.ALLEGRO_CHANNEL_CONF_1
	ChannelConf2 = C.ALLEGRO_CHANNEL_CONF_2
	ChannelConf3 = C.ALLEGRO_CHANNEL_CONF_3
	ChannelConf4 = C.ALLEGRO_CHANNEL_CONF_4
	ChannelConf51 = C.ALLEGRO_CHANNEL_CONF_5_1
	ChannelConf61 = C.ALLEGRO_CHANNEL_CONF_6_1
	ChannelConf71 = C.ALLEGRO_CHANNEL_CONF_7_1
)

const (
	MixerQualityPoint = C.ALLEGRO_MIXER_QUALITY_POINT
	MixerQualityLinear = C.ALLEGRO_MIXER_QUALITY_LINEAR
	MixerQualityCubic = C.ALLEGRO_MIXER_QUALITY_CUBIC
)

const (
	PlaymodeOnce = C.ALLEGRO_PLAYMODE_ONCE
	PlaymodeLoop = C.ALLEGRO_PLAYMODE_LOOP
	PlaymodeBidir = C.ALLEGRO_PLAYMODE_BIDIR
)

const (
	EventAudioRouteChange = C.ALLEGRO_EVENT_AUDIO_ROUTE_CHANGE
	EventAudioInterruption = C.ALLEGRO_EVENT_AUDIO_INTERRUPTION
	EventAudioEndInterruption = C.ALLEGRO_EVENT_AUDIO_END_INTERRUPTION
)

type Mixer C.ALLEGRO_MIXER

type Sample C.ALLEGRO_SAMPLE

type SampleInstance C.ALLEGRO_SAMPLE_INSTANCE

type AudioStream C.ALLEGRO_AUDIO_STREAM

type Voice C.ALLEGRO_VOICE

/********************/
/* Setting up audio */
/********************/

func InstallAudio() bool {
	return bool(C.al_install_audio())
}

func UninstallAudio() {
	C.al_uninstall_audio()
}

func IsAudioInstalled() bool {
	return bool(C.al_is_audio_installed())
}

func ReserveSamples(reserveSamples int32) bool {
	return bool(C.al_reserve_samples(C.int(reserveSamples)))
}

/************************/
/* Misc audio functions */
/************************/

func GetAllegroAudioVersion() uint32 {
	return uint32(C.al_get_allegro_audio_version())
}

func GetAudioDepthSize(depth uint32) uint32 {
	return uint32(C.al_get_audio_depth_size(C.ALLEGRO_AUDIO_DEPTH(depth)))
}

func GetChannelCount(conf uint32) uint32 {
	return uint32(C.al_get_channel_count(C.ALLEGRO_CHANNEL_CONF(conf)))
}

/*******************/
/* Voice functions */
/*******************/

func CreateVoice(freq, depth, chanConf uint32) *Voice {
	return (*Voice)(unsafe.Pointer(C.al_create_voice(C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chanConf))))
}

func (v *Voice) Destroy() {
	C.al_destroy_voice((*C.ALLEGRO_VOICE)(unsafe.Pointer(v)))
}

func (v *Voice) Detach() {
	C.al_detach_voice((*C.ALLEGRO_VOICE)(unsafe.Pointer(v)))
}

func (v *Voice) AttachAudioStream(stream *AudioStream) bool {
	return bool(C.al_attach_audio_stream_to_voice((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(stream)), (*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) AttachMixer(mixer *Mixer) bool {
	return bool(C.al_attach_mixer_to_voice((*C.ALLEGRO_MIXER)(unsafe.Pointer(mixer)), (*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) AttachSampleInstance(spl *SampleInstance) bool {
	return bool(C.al_attach_sample_instance_to_voice((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(spl)), (*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) GetFrequency() uint32 {
	return uint32(C.al_get_voice_frequency((*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) GetChannels() uint32 {
	return uint32(C.al_get_voice_channels((*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) GetDepth() uint32 {
	return uint32(C.al_get_voice_depth((*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) GetPlaying() bool {
	return bool(C.al_get_voice_playing((*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) SetPlaying(val bool) bool {
	return bool(C.al_set_voice_playing((*C.ALLEGRO_VOICE)(unsafe.Pointer(v)), C.bool(val)))
}

func (v *Voice) GetPosition() uint32 {
	return uint32(C.al_get_voice_position((*C.ALLEGRO_VOICE)(unsafe.Pointer(v))))
}

func (v *Voice) SetPosition(val uint32) bool {
	return bool(C.al_set_voice_position((*C.ALLEGRO_VOICE)(unsafe.Pointer(v)), C.uint(val)))
}

/********************/
/* Sample functions */
/********************/

func CreateSample(buf unsafe.Pointer, samples, freq, depth, chanConf uint32, freeBuf bool) *Sample {
	return (*Sample)(unsafe.Pointer(C.al_create_sample(buf, C.uint(samples), C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chanConf), C.bool(freeBuf))))
}

func (s *Sample) Destroy() {
	C.al_destroy_sample((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s)))
}


