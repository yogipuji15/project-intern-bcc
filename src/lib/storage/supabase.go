package storage

import (
	"mime/multipart"
	"os"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type StorageInterface interface {
	UploadFile(file *multipart.FileHeader) (string,error)
}

type storage struct {

}

func Init() StorageInterface {
	return &storage{}
}

func (s *storage) UploadFile(file *multipart.FileHeader) (string,error) {
	link, err := supabasestorageuploader.Upload(os.Getenv("HOST"), os.Getenv("TOKEN"), os.Getenv("STORAGE_NAME"), os.Getenv("STORAGE_PATH"), file)
	if err != nil {
		return link,err
	}
	return link,nil
}