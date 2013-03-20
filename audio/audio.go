package audio

// #cgo LDFLAGS: -lallegro_audio
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_audio.h>
// #include "../events.h"
import "C"

import (
	"github.com/tapir/allegro"
	"unsafe"
)

const (
	DepthInt8     = C.ALLEGRO_AUDIO_DEPTH_INT8
	DepthInt16    = C.ALLEGRO_AUDIO_DEPTH_INT16
	DepthInt24    = C.ALLEGRO_AUDIO_DEPTH_INT24
	DepthFloat32  = C.ALLEGRO_AUDIO_DEPTH_FLOAT32
	DepthUnsigned = C.ALLEGRO_AUDIO_DEPTH_UNSIGNED
	DepthUint8    = C.ALLEGRO_AUDIO_DEPTH_UINT8
	DepthUint16   = C.ALLEGRO_AUDIO_DEPTH_UINT16
	DepthUint24   = C.ALLEGRO_AUDIO_DEPTH_UINT24
)

// const AudioPanNone = C.ALLEGRO_AUDIO_PAN_NONE !??
const AudioPanNone float32 = 1000

const (
	ChannelConf1  = C.ALLEGRO_CHANNEL_CONF_1
	ChannelConf2  = C.ALLEGRO_CHANNEL_CONF_2
	ChannelConf3  = C.ALLEGRO_CHANNEL_CONF_3
	ChannelConf4  = C.ALLEGRO_CHANNEL_CONF_4
	ChannelConf51 = C.ALLEGRO_CHANNEL_CONF_5_1
	ChannelConf61 = C.ALLEGRO_CHANNEL_CONF_6_1
	ChannelConf71 = C.ALLEGRO_CHANNEL_CONF_7_1
)

const (
	MixerQualityPoint  = C.ALLEGRO_MIXER_QUALITY_POINT
	MixerQualityLinear = C.ALLEGRO_MIXER_QUALITY_LINEAR
	MixerQualityCubic  = C.ALLEGRO_MIXER_QUALITY_CUBIC
)

const (
	PlaymodeOnce  = C.ALLEGRO_PLAYMODE_ONCE
	PlaymodeLoop  = C.ALLEGRO_PLAYMODE_LOOP
	PlaymodeBidir = C.ALLEGRO_PLAYMODE_BIDIR
)

const (
	EventRouteChange     = C.ALLEGRO_EVENT_AUDIO_ROUTE_CHANGE
	EventInterruption    = C.ALLEGRO_EVENT_AUDIO_INTERRUPTION
	EventEndInterruption = C.ALLEGRO_EVENT_AUDIO_END_INTERRUPTION
)

const EventRecorderFragment = C.ALLEGRO_EVENT_AUDIO_RECORDER_FRAGMENT

type Mixer C.ALLEGRO_MIXER

type SampleId C.ALLEGRO_SAMPLE_ID

type Sample C.ALLEGRO_SAMPLE

type SampleInstance C.ALLEGRO_SAMPLE_INSTANCE

type Stream C.ALLEGRO_AUDIO_STREAM

type Voice C.ALLEGRO_VOICE

type Recorder C.ALLEGRO_AUDIO_RECORDER

type RecorderEvent struct {
	Type      uint32
	Source    *Recorder
	TimeStamp float64
	internalDescr *C.ALLEGRO_USER_EVENT_DESCRIPTOR
	Buffer unsafe.Pointer
	Samples uint32
}

/********************/
/* Setting up audio */
/********************/

func Install() bool {
	return bool(C.al_install_audio())
}

func Uninstall() {
	C.al_uninstall_audio()
}

func IsInstalled() bool {
	return bool(C.al_is_audio_installed())
}

func ReserveSamples(reserveSamples int32) bool {
	return bool(C.al_reserve_samples(C.int(reserveSamples)))
}

/************************/
/* Misc audio functions */
/************************/

func GetVersion() uint32 {
	return uint32(C.al_get_allegro_audio_version())
}

func GetDepthSize(depth uint32) uint32 {
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

func (v *Voice) AttachStream(stream *Stream) bool {
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

func CreateMixer(freq, depth, chanConf uint32) *Mixer {
	return (*Mixer)(unsafe.Pointer(C.al_create_mixer(C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chanConf))))
}

func (m *Mixer) Destroy() {
	C.al_destroy_mixer((*C.ALLEGRO_MIXER)(unsafe.Pointer(m)))
}

func GetDefaultMixer() *Mixer {
	return (*Mixer)(unsafe.Pointer(C.al_get_default_mixer()))
}

func (m *Mixer) SetDefault() bool {
	return bool(C.al_set_default_mixer((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func RestoreDefaultMixer() bool {
	return bool(C.al_restore_default_mixer())
}

func (m *Mixer) AttachMixer(mixer *Mixer) bool {
	return bool(C.al_attach_mixer_to_mixer((*C.ALLEGRO_MIXER)(unsafe.Pointer(m)), (*C.ALLEGRO_MIXER)(unsafe.Pointer(mixer))))
}

func (m *Mixer) AttachSampleInstance(spl *SampleInstance) bool {
	return bool(C.al_attach_sample_instance_to_mixer((*C.ALLEGRO_SAMPLE_INSTANCE)(unsafe.Pointer(spl)), (*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) AttachStream(stream *Stream) bool {
	return bool(C.al_attach_audio_stream_to_mixer((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(stream)), (*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) GetFrequency() uint32 {
	return uint32(C.al_get_mixer_frequency((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) SetFrequency(freq uint32) bool {
	return bool(C.al_set_mixer_frequency((*C.ALLEGRO_MIXER)(unsafe.Pointer(m)), C.uint(freq)))
}

func (m *Mixer) GetChannels() uint32 {
	return uint32(C.al_get_mixer_channels((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) GetDepth() uint32 {
	return uint32(C.al_get_mixer_depth((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) GetGain() float32 {
	return float32(C.al_get_mixer_gain((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) SetGain(gain float32) bool {
	return bool(C.al_set_mixer_gain((*C.ALLEGRO_MIXER)(unsafe.Pointer(m)), C.float(gain)))
}

func (m *Mixer) GetQuality() uint32 {
	return uint32(C.al_get_mixer_quality((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) SetQuality(newQuality uint32) bool {
	return bool(C.al_set_mixer_quality((*C.ALLEGRO_MIXER)(unsafe.Pointer(m)), C.ALLEGRO_MIXER_QUALITY(newQuality)))
}

func (m *Mixer) GetPlaying() bool {
	return bool(C.al_get_mixer_playing((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) SetPlaying(val bool) bool {
	return bool(C.al_set_mixer_playing((*C.ALLEGRO_MIXER)(unsafe.Pointer(m)), C.bool(val)))
}

func (m *Mixer) GetAttached() bool {
	return bool(C.al_get_mixer_attached((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

func (m *Mixer) Detach() bool {
	return bool(C.al_detach_mixer((*C.ALLEGRO_MIXER)(unsafe.Pointer(m))))
}

/********************/
/* Stream functions */
/********************/

func CreateStream(fragmentCount uint64, fragSample, freq, depth, chanConf uint32) *Stream {
	return (*Stream)(unsafe.Pointer(C.al_create_audio_stream(C.size_t(fragmentCount), C.uint(fragSample), C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chanConf))))
}

func (a *Stream) Destroy() {
	C.al_destroy_audio_stream((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)))
}

func (a *Stream) GetEventSource() *allegro.EventSource {
	return (*allegro.EventSource)(unsafe.Pointer(C.al_get_audio_stream_event_source((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)))))
}

func (a *Stream) Drain() {
	C.al_drain_audio_stream((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)))
}

func (a *Stream) Rewind() bool {
	return bool(C.al_rewind_audio_stream((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetFrequency() uint32 {
	return uint32(C.al_get_audio_stream_frequency((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetChannels() uint32 {
	return uint32(C.al_get_audio_stream_channels((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetDepth() uint32 {
	return uint32(C.al_get_audio_stream_depth((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetLength() uint32 {
	return uint32(C.al_get_audio_stream_length((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetSpeed() float32 {
	return float32(C.al_get_audio_stream_speed((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetGain() float32 {
	return float32(C.al_get_audio_stream_gain((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) SetGain(val float32) bool {
	return bool(C.al_set_audio_stream_gain((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)), C.float(val)))
}

func (a *Stream) GetPan() float32 {
	return float32(C.al_get_audio_stream_pan((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) SetPan(val float32) bool {
	return bool(C.al_set_audio_stream_pan((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)), C.float(val)))
}

func (a *Stream) GetPlaying() bool {
	return bool(C.al_get_audio_stream_playing((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) SetPlaying(val bool) bool {
	return bool(C.al_set_audio_stream_playing((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)), C.bool(val)))
}

func (a *Stream) GetPlaymode() uint32 {
	return uint32(C.al_get_audio_stream_playmode((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) SetPlaymode(val uint32) bool {
	return bool(C.al_set_audio_stream_playmode((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)), C.ALLEGRO_PLAYMODE(val)))
}

func (a *Stream) GetAttached() bool {
	return bool(C.al_get_audio_stream_attached((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) Detach() bool {
	return bool(C.al_detach_audio_stream((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetFragment() unsafe.Pointer {
	return unsafe.Pointer(C.al_get_audio_stream_fragment((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) SetFragment(val unsafe.Pointer) bool {
	return bool(C.al_set_audio_stream_fragment((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)), val))
}

func (a *Stream) GetFragments() uint32 {
	return uint32(C.al_get_audio_stream_fragments((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetAvailableFragments() uint32 {
	return uint32(C.al_get_available_audio_stream_fragments((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) SeekSecs(time float64) bool {
	return bool(C.al_seek_audio_stream_secs((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)), C.double(time)))
}

func (a *Stream) GetPositionSecs() float64 {
	return float64(C.al_get_audio_stream_position_secs((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) GetLengthSecs() float64 {
	return float64(C.al_get_audio_stream_length_secs((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a))))
}

func (a *Stream) SetLoopSecs(start, end float64) bool {
	return bool(C.al_set_audio_stream_loop_secs((*C.ALLEGRO_AUDIO_STREAM)(unsafe.Pointer(a)), C.double(start), C.double(end)))
}

/*****************/
/* Audio file IO */
/*****************/

func LoadSample(filename string) *Sample {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return (*Sample)(unsafe.Pointer(C.al_load_sample(f)))
}

func LoadSampleF(fp *allegro.File, ident string) *Sample {
	i := C.CString(ident)
	defer C.free(unsafe.Pointer(i))
	return (*Sample)(unsafe.Pointer(C.al_load_sample_f((*C.ALLEGRO_FILE)(unsafe.Pointer(fp)), i)))
}

func LoadStream(filename string, bufferCount uint64, samples uint32) *Stream {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return (*Stream)(unsafe.Pointer(C.al_load_audio_stream(f, C.size_t(bufferCount), C.uint(samples))))
}

func LoadStreamF(fp *allegro.File, ident string, bufferCount uint64, samples uint32) *Stream {
	i := C.CString(ident)
	defer C.free(unsafe.Pointer(i))
	return (*Stream)(unsafe.Pointer(C.al_load_audio_stream_f((*C.ALLEGRO_FILE)(fp), i, C.size_t(bufferCount), C.uint(samples))))
}

func (s *Sample) Save(filename string) bool {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return bool(C.al_save_sample(f, (*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s))))
}

func (s *Sample) SaveF(fp *allegro.File, ident string) bool {
	i := C.CString(ident)
	defer C.free(unsafe.Pointer(i))
	return bool(C.al_save_sample_f((*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s)), i, (*C.ALLEGRO_SAMPLE)(unsafe.Pointer(s))))
}

/****************/
/* Audio events */
/****************/

func GetEventSource() *allegro.EventSource {
	return (*allegro.EventSource)(unsafe.Pointer(C.al_get_audio_event_source()))
}

/*******************/
/* Audio recording */
/*******************/

func CreateRecorder(fragmentCount uint64, samples, frequency, depth, chanConf uint32) *Recorder {
	return (*Recorder)(unsafe.Pointer(C.al_create_audio_recorder(C.size_t(fragmentCount), C.uint(samples), C.uint(frequency), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chanConf))))
}

func (a *Recorder) Destroy() {
	C.al_destroy_audio_recorder((*C.ALLEGRO_AUDIO_RECORDER)(unsafe.Pointer(a)))
}

func (a *Recorder) Stop() {
	C.al_stop_audio_recorder((*C.ALLEGRO_AUDIO_RECORDER)(unsafe.Pointer(a)))
}

func (a *Recorder) IsRecording() bool {
	return bool(C.al_is_audio_recorder_recording((*C.ALLEGRO_AUDIO_RECORDER)(unsafe.Pointer(a))))
}

//func (e *allegro.Event) GetEvent() *RecorderEvent {
//	return (*RecorderEvent)(unsafe.Pointer(C.al_get_audio_recorder_event((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
//}

func (a *Recorder) GetEventSource() *allegro.EventSource {
	return (*allegro.EventSource)(unsafe.Pointer(C.al_get_audio_recorder_event_source((*C.ALLEGRO_AUDIO_RECORDER)(unsafe.Pointer(a)))))
}
