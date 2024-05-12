package configs

type UploadMediaData struct {
	Key      string `bson:"key" json:"key" validate:"required"`
	FileName string `bson:"file_name" json:"file_name" validate:"required"`
	FileData []byte `bson:"file_data" json:"file_data" validate:"required"`
	FileSize int64  `bson:"file_size" json:"file_size" validate:"required"`
}

type UploadMediaRequest struct {
	FileName string `bson:"file_name" json:"file_name" validate:"required"`
	FileSize int64  `bson:"file_size" json:"file_size" validate:"required"`
}

type DeleteObjectReq struct {
	FileName string `bson:"file_name" json:"file_name" validate:"required"`
}
