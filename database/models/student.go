package models

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Name       string `json:"name"`
	Username   string `json:"username"`
	LevelOne   bool   `json:"level_0"`
	LevelTwo   bool   `json:"level_1"`
	LevelThree bool   `json:"level_2"`
	LevelFour  bool   `json:"level_3"`
	LevelFive  bool   `json:"level_4"`
	LevelSix   bool   `json:"level_5"`
	LevelSeven bool   `json:"level_6"`
}
