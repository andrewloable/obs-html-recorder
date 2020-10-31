package scene

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/andrewloable/obs-html-recorder/config"
)

// GenerateDefaultScene - Generate the default scene file in obs
func GenerateDefaultScene() (bool, error) {
	// get all files in the scene folder then delete
	scenePath := config.AppConfig.ObsSettingsPath + "/basic/scenes"
	files, err := ioutil.ReadDir(scenePath)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		filePath := scenePath + "/" + file.Name()
		err = os.Remove(filePath)
		if err != nil {
			return false, err
		}
	}
	// write scene file
	sceneFilePath := config.AppConfig.ObsSettingsPath + "/" + GenerateSceneFilePath()
	f, err := os.Create(sceneFilePath)
	log.Println("create file " + sceneFilePath)
	if err != nil {
		log.Println("error: scene create file")
		return false, err
	}
	defer f.Close()

	_, err = f.WriteString(GenerateScene())
	if err != nil {
		log.Println("error: write scene settings")
		return false, err
	}
	return true, nil
}

// GenerateSceneFilePath - Generate a file name based on the profile
func GenerateSceneFilePath() string {
	return "basic/scenes/default.json"
}

// GenerateScene - Returns scene settings string
func GenerateScene() string {
	return fmt.Sprintf(`
	{
		"AuxAudioDevice1": {
			"balance": 0.5,
			"deinterlace_field_order": 0,
			"deinterlace_mode": 0,
			"enabled": true,
			"flags": 0,
			"hotkeys": {
				"libobs.mute": [],
				"libobs.push-to-mute": [],
				"libobs.push-to-talk": [],
				"libobs.unmute": []
			},
			"id": "wasapi_input_capture",
			"mixers": 255,
			"monitoring_type": 0,
			"muted": false,
			"name": "Mic/Aux",
			"prev_ver": 436207618,
			"private_settings": {},
			"push-to-mute": false,
			"push-to-mute-delay": 0,
			"push-to-talk": false,
			"push-to-talk-delay": 0,
			"settings": {
				"device_id": "default"
			},
			"sync": 0,
			"versioned_id": "wasapi_input_capture",
			"volume": 1.0
		},
		"DesktopAudioDevice1": {
			"balance": 0.5,
			"deinterlace_field_order": 0,
			"deinterlace_mode": 0,
			"enabled": true,
			"flags": 0,
			"hotkeys": {
				"libobs.mute": [],
				"libobs.push-to-mute": [],
				"libobs.push-to-talk": [],
				"libobs.unmute": []
			},
			"id": "wasapi_output_capture",
			"mixers": 255,
			"monitoring_type": 0,
			"muted": true,
			"name": "Desktop Audio",
			"prev_ver": 436207618,
			"private_settings": {},
			"push-to-mute": false,
			"push-to-mute-delay": 0,
			"push-to-talk": false,
			"push-to-talk-delay": 0,
			"settings": {
				"device_id": "default"
			},
			"sync": 0,
			"versioned_id": "wasapi_output_capture",
			"volume": 1.0
		},
		"current_program_scene": "default",
		"current_scene": "default",
		"current_transition": "Fade",
		"groups": [],
		"modules": {
			"auto-scene-switcher": {
				"active": false,
				"interval": 300,
				"non_matching_scene": "",
				"switch_if_not_matching": false,
				"switches": []
			},
			"captions": {
				"enabled": false,
				"lang_id": 1033,
				"provider": "mssapi",
				"source": ""
			},
			"output-timer": {
				"autoStartRecordTimer": false,
				"autoStartStreamTimer": false,
				"pauseRecordTimer": true,
				"recordTimerHours": 0,
				"recordTimerMinutes": 0,
				"recordTimerSeconds": 30,
				"streamTimerHours": 0,
				"streamTimerMinutes": 0,
				"streamTimerSeconds": 30
			},
			"scripts-tool": []
		},
		"name": "Untitled",
		"preview_locked": false,
		"quick_transitions": [
			{
				"duration": 300,
				"fade_to_black": false,
				"hotkeys": [],
				"id": 1,
				"name": "Cut"
			},
			{
				"duration": 300,
				"fade_to_black": false,
				"hotkeys": [],
				"id": 2,
				"name": "Fade"
			},
			{
				"duration": 300,
				"fade_to_black": true,
				"hotkeys": [],
				"id": 3,
				"name": "Fade"
			}
		],
		"saved_projectors": [],
		"scaling_enabled": false,
		"scaling_level": 0,
		"scaling_off_x": 0.0,
		"scaling_off_y": 0.0,
		"scene_order": [
			{
				"name": "default"
			}
		],
		"sources": [
			{
				"balance": 0.5,
				"deinterlace_field_order": 0,
				"deinterlace_mode": 0,
				"enabled": true,
				"flags": 0,
				"hotkeys": {
					"libobs.mute": [],
					"libobs.push-to-mute": [],
					"libobs.push-to-talk": [],
					"libobs.unmute": []
				},
				"id": "browser_source",
				"mixers": 255,
				"monitoring_type": 0,
				"muted": false,
				"name": "Browser",
				"prev_ver": 436207618,
				"private_settings": {},
				"push-to-mute": false,
				"push-to-mute-delay": 0,
				"push-to-talk": false,
				"push-to-talk-delay": 0,
				"settings": {
					"url": "https://loable.tech"
				},
				"sync": 0,
				"versioned_id": "browser_source",
				"volume": 1.0
			},
			{
				"balance": 0.5,
				"deinterlace_field_order": 0,
				"deinterlace_mode": 0,
				"enabled": true,
				"flags": 0,
				"hotkeys": {
					"OBSBasic.SelectScene": [],
					"libobs.hide_scene_item.Browser": [],
					"libobs.show_scene_item.Browser": []
				},
				"id": "scene",
				"mixers": 0,
				"monitoring_type": 0,
				"muted": false,
				"name": "default",
				"prev_ver": 436207618,
				"private_settings": {},
				"push-to-mute": false,
				"push-to-mute-delay": 0,
				"push-to-talk": false,
				"push-to-talk-delay": 0,
				"settings": {
					"custom_size": false,
					"id_counter": 1,
					"items": [
						{
							"align": 5,
							"bounds": {
								"x": 0.0,
								"y": 0.0
							},
							"bounds_align": 0,
							"bounds_type": 0,
							"crop_bottom": 0,
							"crop_left": 0,
							"crop_right": 0,
							"crop_top": 0,
							"group_item_backup": false,
							"id": 1,
							"locked": false,
							"name": "Browser",
							"pos": {
								"x": 0.0,
								"y": 0.0
							},
							"private_settings": {},
							"rot": 0.0,
							"scale": {
								"x": 1.0,
								"y": 1.0
							},
							"scale_filter": "disable",
							"visible": true
						}
					]
				},
				"sync": 0,
				"versioned_id": "scene",
				"volume": 1.0
			}
		],
		"transition_duration": 300,
		"transitions": []
	}
	`)
}
