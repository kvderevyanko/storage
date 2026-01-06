package storage

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kvderevyanko/storage/internal/file"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	s.Files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("file %v not found", fileId))
	}

	return foundFile, nil
}
