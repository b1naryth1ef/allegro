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

type SampleId C.ALLEGRO_SAMPLE_ID

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

func (s *Sample) Play(gain, pan, speed float32, loop uint32) (bool, *SampleId) {
	retId := new(C.ALLEGRO_SAMPLE_ID)
	r := bool(C.al_play_sample((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s)), C.float(gain), C.float(pan), C.float(speed), C.ALLEGRO_PLAYMODE(loop), retId))
	return r, (*SampleId)(unsafe.Pointer(retId))
}

func (s *SampleId) Stop() {
	C.al_stop_sample((*C.ALLEGRO_SAMPLE_ID)(unsafe.Pointer(s)))
}

func StopSamples() {
	C.al_stop_samples()
}

func (s *Sample) GetChannels() uint32 {
	return uint32(C.al_get_sample_channels((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s))))
}

func (s *Sample) GetDepth() uint32 {
	return uint32(C.al_get_sample_depth((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s))))
}

func (s *Sample) GetFrequency() uint32 {
	return uint32(C.al_get_sample_frequency((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s))))
}

func (s *Sample) GetLength() uint32 {
	return uint32(C.al_get_sample_length((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s))))
}

func (s *Sample) GetData() unsafe.Pointer {
	return unsafe.Pointer(C.al_get_sample_data((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s))))
}

/*****************************/
/* Sample instance functions */
/*****************************/

func CreateSampleInstance(sampleData *Sample) *SampleInstance {
	return (*SampleInstance)(unsafe.Pointer(C.al_create_sample_instance((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(sampleData)))))
}

func (s *SampleInstance) Destroy() {
	C.al_destroy_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)))
}

func (s *SampleInstance) Play() bool {
	return bool(C.al_play_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) Stop() bool {
	return bool(C.al_stop_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) GetChannels() uint32 {
	return uint32(C.al_get_sample_instance_channels((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) GetDepth() uint32 {
	return uint32(C.al_get_sample_instance_depth((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) GetFrequency() uint32 {
	return uint32(C.al_get_sample_instance_frequency((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) GetLength() uint32 {
	return uint32(C.al_get_sample_instance_length((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) SetLength(val uint32) bool {
	return bool(C.al_set_sample_instance_length((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), C.uint(val)))
}

func (s *SampleInstance) GetPosition() uint32 {
	return uint32(C.al_get_sample_instance_position((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) SetPosition(val uint32) bool {
	return bool(C.al_set_sample_instance_position((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), C.uint(val)))
}

func (s *SampleInstance) GetSpeed() float32 {
	return float32(C.al_get_sample_instance_speed((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) SetSpeed(val float32) bool {
	return bool(C.al_set_sample_instance_speed((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), C.float(val)))
}

func (s *SampleInstance) GetGain() float32 {
	return float32(C.al_get_sample_instance_gain((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) SetGain(val float32) bool {
	return bool(C.al_set_sample_instance_gain((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), C.float(val)))
}

func (s *SampleInstance) GetPan() float32 {
	return float32(C.al_get_sample_instance_pan((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) SetPan(val float32) bool {
	return bool(C.al_set_sample_instance_pan((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), C.float(val)))
}

func (s *SampleInstance) GetTime() float32 {
	return float32(C.al_get_sample_instance_time((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) GetPlaymode() uint32 {
	return uint32(C.al_get_sample_instance_playmode((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) SetPlaymode(val uint32) bool {
	return bool(C.al_set_sample_instance_playmode((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), C.ALLEGRO_PLAYMODE(val)))
}

func (s *SampleInstance) GetPlaying() bool {
	return bool(C.al_get_sample_instance_playing((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) SetPlaying(val bool) bool {
	return bool(C.al_set_sample_instance_playing((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), C.bool(val)))
}

func (s *SampleInstance) GetAttached() bool {
	return bool(C.al_get_sample_instance_attached((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) Detach() bool {
	return bool(C.al_detach_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s))))
}

func (s *SampleInstance) GetSample() *Sample {
	return (*Sample)(unsafe.Pointer(C.al_get_sample((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)))))
}

func (s *SampleInstance) SetSample(data *Sample) bool {
	return bool(C.al_set_sample((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(s)), (*C.ALLEGRO_SAMPLE)(unsafe.Pointer(data))))
}

/*******************/
/* Mixer functions */
/*******************/


