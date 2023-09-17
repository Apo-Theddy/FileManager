package models

import "time"

type File struct {
	FileID       uint       `gorm:"column:FileID; type:uint;not null;autoIncrement;primaryKey;" json:"file_id"`
	UniqueUUID   string     `gorm:"column:UniqueUUID; type:varchar(100);unique;not null;" json:"unique_uuid"`
	OriginalName string     `gorm:"column:OriginalName; type:varchar(150);not null;" json:"original_name"`
	DirUploaded  string     `gorm:"column:DirUploaded; type:varchar(150);default:'uploads/';null;" json:"dir_uploaded"`
	UploadDate   time.Time  `gorm:"column:UploadDate; type:date;default:CURRENT_TIMESTAMP;null;" json:"upload_date"`
	DeleteDate   *time.Time `gorm:"column:DeleteDate; type:date;null;default:null;" json:"delete_date"`
}

func (File) TableName() string {
	return "Files"
}
