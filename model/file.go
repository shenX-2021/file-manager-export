package model

type FileModel struct {
	ID              int `gorm:"primarykey"`
	FileName        string
	FileHash        string
	StartHash       string
	EndHash         string
	FilePath        string
	Size            int
	Status          int
	CheckStatus     int
	OutsideDownload int
	GmtCreated      int
	GmtModified     int
}

func (FileModel) TableName() string {
	return "tb_file"
}
