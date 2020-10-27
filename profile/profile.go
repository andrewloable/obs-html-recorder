package profile

import (
	"fmt"
	"strconv"
)

// Profile struct
type Profile struct {
	Width  int
	Height int
}

// GenerateProfileFilePath - Generate a file name based on the profile
func (p Profile) GenerateProfileFilePath() string {
	return "basic/profiles/" + p.ResolutionString() + "/basic.ini"
}

// ResolutionString - Get width x height string
func (p Profile) ResolutionString() string {
	return strconv.Itoa(p.Width) + "x" + strconv.Itoa(p.Height)
}

func (p Profile) widthString() string {
	return strconv.Itoa(p.Width)
}

func (p Profile) heightString() string {
	return strconv.Itoa(p.Height)
}

// GenerateSettings - Generate settings string
func (p Profile) GenerateSettings() string {
	return fmt.Sprintf(`
	[General]
	Name=%s

	[Twitch]
	AddonChoice=3
	
	[Video]
	BaseCX=%s
	BaseCY=%s
	OutputCX=%s
	OutputCY=%s
	
	[Panels]
	CookieId=F2EB79A2DEF4AE3A
	
	[Output]
	Mode=Simple
	
	[AdvOut]
	RescaleRes=%s
	TrackIndex=1
	RecType=Standard
	RecRescaleRes=%s
	RecTracks=1
	FLVTrack=1
	FFOutputToFile=true
	FFFormat=
	FFFormatMimeType=
	FFRescaleRes=%s
	FFVEncoderId=0
	FFVEncoder=
	FFAEncoderId=0
	FFAEncoder=
	FFAudioMixes=1
	
	[SimpleOutput]
	FileNameWithoutSpace=true
	RecFormat=mp4
	RecQuality=HQ

	`, p.ResolutionString(),
		p.widthString(),
		p.heightString(),
		p.widthString(),
		p.heightString(),
		p.ResolutionString(),
		p.ResolutionString(),
		p.ResolutionString())
}
