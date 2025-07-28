package models

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Level_0  bool   `json:"level_0"`
	Level_1  bool   `json:"level_1"`
	Level_2  bool   `json:"level_2"`
	Level_3  bool   `json:"level_3"`
	Level_4  bool   `json:"level_4"`
	Level_5  bool   `json:"level_5"`
	Level_6  bool   `json:"level_6"`
	Level_7  bool   `json:"level_7"`
}
