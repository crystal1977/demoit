package emulation

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// ScreenOrientation screen orientation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-ScreenOrientation
type ScreenOrientation struct {
	Type  OrientationType `json:"type"`  // Orientation type.
	Angle int64           `json:"angle"` // Orientation angle.
}

// MediaFeature [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-MediaFeature
type MediaFeature struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// VirtualTimePolicy advance: If the scheduler runs out of immediate work,
// the virtual time base may fast forward to allow the next delayed task (if
// any) to run; pause: The virtual time base may not advance;
// pauseIfNetworkFetchesPending: The virtual time base may not advance if there
// are any pending resource fetches.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-VirtualTimePolicy
type VirtualTimePolicy string

// String returns the VirtualTimePolicy as string value.
func (t VirtualTimePolicy) String() string {
	return string(t)
}

// VirtualTimePolicy values.
const (
	VirtualTimePolicyAdvance                      VirtualTimePolicy = "advance"
	VirtualTimePolicyPause                        VirtualTimePolicy = "pause"
	VirtualTimePolicyPauseIfNetworkFetchesPending VirtualTimePolicy = "pauseIfNetworkFetchesPending"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t VirtualTimePolicy) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t VirtualTimePolicy) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *VirtualTimePolicy) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch VirtualTimePolicy(in.String()) {
	case VirtualTimePolicyAdvance:
		*t = VirtualTimePolicyAdvance
	case VirtualTimePolicyPause:
		*t = VirtualTimePolicyPause
	case VirtualTimePolicyPauseIfNetworkFetchesPending:
		*t = VirtualTimePolicyPauseIfNetworkFetchesPending

	default:
		in.AddError(errors.New("unknown VirtualTimePolicy value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *VirtualTimePolicy) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// OrientationType orientation type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-ScreenOrientation
type OrientationType string

// String returns the OrientationType as string value.
func (t OrientationType) String() string {
	return string(t)
}

// OrientationType values.
const (
	OrientationTypePortraitPrimary    OrientationType = "portraitPrimary"
	OrientationTypePortraitSecondary  OrientationType = "portraitSecondary"
	OrientationTypeLandscapePrimary   OrientationType = "landscapePrimary"
	OrientationTypeLandscapeSecondary OrientationType = "landscapeSecondary"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t OrientationType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t OrientationType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *OrientationType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch OrientationType(in.String()) {
	case OrientationTypePortraitPrimary:
		*t = OrientationTypePortraitPrimary
	case OrientationTypePortraitSecondary:
		*t = OrientationTypePortraitSecondary
	case OrientationTypeLandscapePrimary:
		*t = OrientationTypeLandscapePrimary
	case OrientationTypeLandscapeSecondary:
		*t = OrientationTypeLandscapeSecondary

	default:
		in.AddError(errors.New("unknown OrientationType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *OrientationType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SetEmitTouchEventsForMouseConfiguration touch/gesture events
// configuration. Default: current platform.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setEmitTouchEventsForMouse
type SetEmitTouchEventsForMouseConfiguration string

// String returns the SetEmitTouchEventsForMouseConfiguration as string value.
func (t SetEmitTouchEventsForMouseConfiguration) String() string {
	return string(t)
}

// SetEmitTouchEventsForMouseConfiguration values.
const (
	SetEmitTouchEventsForMouseConfigurationMobile  SetEmitTouchEventsForMouseConfiguration = "mobile"
	SetEmitTouchEventsForMouseConfigurationDesktop SetEmitTouchEventsForMouseConfiguration = "desktop"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetEmitTouchEventsForMouseConfiguration) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetEmitTouchEventsForMouseConfiguration) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetEmitTouchEventsForMouseConfiguration) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SetEmitTouchEventsForMouseConfiguration(in.String()) {
	case SetEmitTouchEventsForMouseConfigurationMobile:
		*t = SetEmitTouchEventsForMouseConfigurationMobile
	case SetEmitTouchEventsForMouseConfigurationDesktop:
		*t = SetEmitTouchEventsForMouseConfigurationDesktop

	default:
		in.AddError(errors.New("unknown SetEmitTouchEventsForMouseConfiguration value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetEmitTouchEventsForMouseConfiguration) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SetEmulatedVisionDeficiencyType vision deficiency to emulate.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setEmulatedVisionDeficiency
type SetEmulatedVisionDeficiencyType string

// String returns the SetEmulatedVisionDeficiencyType as string value.
func (t SetEmulatedVisionDeficiencyType) String() string {
	return string(t)
}

// SetEmulatedVisionDeficiencyType values.
const (
	SetEmulatedVisionDeficiencyTypeNone          SetEmulatedVisionDeficiencyType = "none"
	SetEmulatedVisionDeficiencyTypeAchromatopsia SetEmulatedVisionDeficiencyType = "achromatopsia"
	SetEmulatedVisionDeficiencyTypeBlurredVision SetEmulatedVisionDeficiencyType = "blurredVision"
	SetEmulatedVisionDeficiencyTypeDeuteranopia  SetEmulatedVisionDeficiencyType = "deuteranopia"
	SetEmulatedVisionDeficiencyTypeProtanopia    SetEmulatedVisionDeficiencyType = "protanopia"
	SetEmulatedVisionDeficiencyTypeTritanopia    SetEmulatedVisionDeficiencyType = "tritanopia"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetEmulatedVisionDeficiencyType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetEmulatedVisionDeficiencyType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetEmulatedVisionDeficiencyType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SetEmulatedVisionDeficiencyType(in.String()) {
	case SetEmulatedVisionDeficiencyTypeNone:
		*t = SetEmulatedVisionDeficiencyTypeNone
	case SetEmulatedVisionDeficiencyTypeAchromatopsia:
		*t = SetEmulatedVisionDeficiencyTypeAchromatopsia
	case SetEmulatedVisionDeficiencyTypeBlurredVision:
		*t = SetEmulatedVisionDeficiencyTypeBlurredVision
	case SetEmulatedVisionDeficiencyTypeDeuteranopia:
		*t = SetEmulatedVisionDeficiencyTypeDeuteranopia
	case SetEmulatedVisionDeficiencyTypeProtanopia:
		*t = SetEmulatedVisionDeficiencyTypeProtanopia
	case SetEmulatedVisionDeficiencyTypeTritanopia:
		*t = SetEmulatedVisionDeficiencyTypeTritanopia

	default:
		in.AddError(errors.New("unknown SetEmulatedVisionDeficiencyType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetEmulatedVisionDeficiencyType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
