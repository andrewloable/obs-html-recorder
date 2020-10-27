package obs

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/andrewloable/obs-html-recorder/config"
	"github.com/andrewloable/obs-html-recorder/profile"
	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/mitchellh/go-ps"
)

// RecordHTML - Record html to video for x seconds
func RecordHTML(client obsws.Client, url string, seconds int, prof profile.Profile) (string, error) {
	// set browser source
	log.Println("add browser source")
	obsws.NewSetBrowserSourcePropertiesRequest("Browser", false, "", url, "", prof.Width, prof.Height, 60, false, true).SendReceive(client)

	// set recording
	outputPath, _ := filepath.Abs("./recordings")
	_, err := os.Stat(outputPath)
	if err == nil {
		// output path already exists
	} else {
		if os.IsNotExist(err) {
			os.MkdirAll(outputPath, os.ModePerm)
		} else {
			// other error
			log.Println("error: other error")
			return "", err
		}
	}

	log.Println("set recording output", outputPath)
	obsws.NewSetRecordingFolderRequest(outputPath).SendReceive(client)
	time.Sleep(time.Second * 1)

	// start recording
	log.Println("start recording")
	obsws.NewStartRecordingRequest().SendReceive(client)

	// wait seconds
	log.Println("waiting ", seconds)
	time.Sleep(time.Second * time.Duration(seconds))

	// stop recording
	log.Println("stop recording")
	obsws.NewStopRecordingRequest().SendReceive(client)
	time.Sleep(time.Second * 1)

	log.Println("reset source")
	obsws.NewSetBrowserSourcePropertiesRequest("Browser", false, "", "https://obsproject.com/browser-source", "", prof.Width, prof.Height, 60, false, true).SendReceive(client)
	return "", nil
}

// InitiateObsRecorder - Run Obs with a specific profile
func InitiateObsRecorder(prof profile.Profile) (obsws.Client, error) {
	err := TerminateObs()
	log.Println("terminate obs")
	if err != nil {
		log.Println("error: terminate obs")
		return obsws.Client{}, err
	}

	profilePath := config.AppConfig.ObsSettingsPath + "/" + prof.GenerateProfileFilePath()
	log.Println("profile path " + profilePath)
	_, err = os.Stat(profilePath)
	if err == nil {
		// profile already exists
		log.Println("profile already exists, overwriting")
		os.Remove(profilePath)
		f, err := os.Create(profilePath)
		log.Println("create file " + profilePath)
		if err != nil {
			log.Println("error: create file")
			return obsws.Client{}, err
		}
		defer f.Close()

		_, err = f.WriteString(prof.GenerateSettings())
		if err != nil {
			log.Println("error: write settings")
			return obsws.Client{}, err
		}
	} else {
		if os.IsNotExist(err) {
			log.Println("profile does not exist")
			// profile does not exist, generate a new one
			profileDir := filepath.Dir(profilePath)
			os.MkdirAll(profileDir, os.ModePerm)
			log.Println("create dir " + profileDir)
			f, err := os.Create(profilePath)
			log.Println("create file " + profilePath)
			if err != nil {
				log.Println("error: create file")
				return obsws.Client{}, err
			}
			defer f.Close()

			_, err = f.WriteString(prof.GenerateSettings())
			if err != nil {
				log.Println("error: write settings")
				return obsws.Client{}, err
			}
		} else {
			// other error
			log.Println("error: other error")
			return obsws.Client{}, err
		}
	}

	log.Println("run obs")
	err = RunObs()
	if err != nil {
		log.Println("error: runobs")
		return obsws.Client{}, err
	}

	time.Sleep(time.Second * 5)

	log.Println("obs client initiate")
	client := obsws.Client{
		Host: "localhost",
		Port: 4444,
	}
	if err := client.Connect(); err != nil {
		log.Println("connection", err)
	}
	obsws.SetReceiveTimeout(time.Second * 5)

	log.Println("set obs profile")
	err = SetObsProfile(client, prof.ResolutionString())
	if err != nil {
		log.Println("error: set obs profile")
		return obsws.Client{}, err
	}
	return client, nil
}

// SetObsProfile - Set the obs profile
func SetObsProfile(client obsws.Client, profileName string) error {
	req := obsws.NewSetCurrentProfileRequest(profileName)
	_, err := req.SendReceive(client)
	if err != nil {
		return err
	}
	return nil
}

// CheckObsWebsocket - Check if OBS is running with websockets, returns true if connected
func CheckObsWebsocket(client obsws.Client) (bool, error) {
	req := obsws.NewSetCurrentProfileRequest("1920x1080")
	_, err := req.SendReceive(client)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getObsInstancePid() (int, error) {
	if runtime.GOOS == "windows" {
		proclist, err := ps.Processes()
		if err != nil {
			return 0, err
		}
		for x := range proclist {
			process := proclist[x]
			if process.Executable() == "obs64.exe" {
				return process.Pid(), nil
			}
		}
	}
	return 0, nil
}

// CheckObsIsRunning - Returns true if OBS is currently running
func CheckObsIsRunning() (bool, error) {
	val, err := getObsInstancePid()
	if err != nil {
		return false, err
	}
	if val > 0 {
		return true, nil
	}
	return false, nil
}

// TerminateObs - Closes any OBS instances currently running
func TerminateObs() error {
	val, err := getObsInstancePid()
	if err != nil {
		return err
	}
	if val > 0 {
		if runtime.GOOS == "windows" {
			cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(val))
			err := cmd.Run()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RunObs - Runs a new instance of OBS
func RunObs() error {
	cmd := exec.Command(config.AppConfig.ObsExePath)
	cmd.Dir = filepath.Dir(config.AppConfig.ObsExePath)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
