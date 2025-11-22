package services

import (
	"fmt"
	"os"
	"path/filepath"

	"tugas8/app/model"
	"tugas8/app/repository"
	"tugas8/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FileService interface {
	UploadFile(c *fiber.Ctx) error
	GetAllFiles(c *fiber.Ctx) error
	GetFileByID(c *fiber.Ctx) error
	DeleteFile(c *fiber.Ctx) error
}

type fileService struct {
	repo       repository.FileRepository
	uploadPath string
}

func NewFileService(repo repository.FileRepository, uploadPath string) FileService {
	return &fileService{
		repo:       repo,
		uploadPath: uploadPath,
	}
}

// UploadFile godoc
// @Summary      Upload file
// @Description  Upload file (PDF, JPG, JPEG, PNG) maksimal 10MB
// @Tags         Upload
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "File yang akan diupload"
// @Success      201 {object} map[string]interface{} "File uploaded successfully"
// @Failure      400 {object} map[string]interface{} "Bad request"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /upload [post]
func (s *fileService) UploadFile(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No file uploaded",
			"error":   err.Error(),
		})
	}

	// Validasi ukuran (max 10MB)
	if fileHeader.Size > 10*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "File size exceeds 10MB",
		})
	}

	// Validasi tipe file
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"application/pdf": true,
	}
	contentType := fileHeader.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "File type not allowed. Only JPG, PNG, PDF",
		})
	}

	// Generate nama unik
	ext := filepath.Ext(fileHeader.Filename)
	newFileName := uuid.New().String() + ext
	filePath := filepath.Join(s.uploadPath, newFileName)

	// Buat folder uploads kalau belum ada
	if err := os.MkdirAll(s.uploadPath, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create upload directory",
			"error":   err.Error(),
		})
	}

	// Simpan file ke disk
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to open uploaded file",
			"error":   err.Error(),
		})
	}
	defer file.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file",
			"error":   err.Error(),
		})
	}
	defer out.Close()

	if _, err = out.ReadFrom(file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to write file",
			"error":   err.Error(),
		})
	}

	// Simpan metadata ke MongoDB
	fileModel := &model.File{
		FileName:     newFileName,
		OriginalName: fileHeader.Filename,
		FilePath:     filePath,
		FileSize:     fileHeader.Size,
		FileType:     contentType,
	}

	if err := s.repo.Create(fileModel); err != nil {
		os.Remove(filePath) // rollback
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file metadata",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "File uploaded successfully",
		"data":    s.toFileResponse(fileModel),
	})
}

// GetAllFiles godoc
// @Summary      Ambil semua file yang diupload
// @Tags         Upload
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Router       /upload [get]
func (s *fileService) GetAllFiles(c *fiber.Ctx) error {
	files, err := s.repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get files",
			"error":   err.Error(),
		})
	}

	var responses []model.FileResponse
	for _, f := range files {
		responses = append(responses, *s.toFileResponse(&f))
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Files retrieved successfully",
		"data":    responses,
	})
}

// GetFileByID godoc
// @Summary      Ambil file berdasarkan ID
// @Tags         Upload
// @Security     BearerAuth
// @Produce      json
// @Param        id path string true "File ID"
// @Success      200 {object} map[string]interface{}
// @Failure      404 {object} map[string]interface{}
// @Router       /upload/{id} [get]
func (s *fileService) GetFileByID(c *fiber.Ctx) error {
	id := c.Params("id")
	file, err := s.repo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "File not found",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "File retrieved successfully",
		"data":    s.toFileResponse(file),
	})
}

// DeleteFile godoc
// @Summary      Hapus file
// @Tags         Upload
// @Security     BearerAuth
// @Param        id path string true "File ID"
// @Success      200 {object} map[string]interface{}
// @Failure      404 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /upload/{id} [delete]
func (s *fileService) DeleteFile(c *fiber.Ctx) error {
	id := c.Params("id")
	file, err := s.repo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "File not found",
		})
	}

	// Hapus file fisik
	if err := os.Remove(file.FilePath); err != nil {
		fmt.Println("Warning: Failed to delete physical file:", err)
	}

	// Hapus dari database
	if err := s.repo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete file record",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "File deleted successfully",
	})
}

// Helper: convert model.File → model.FileResponse
func (s *fileService) toFileResponse(file *model.File) *model.FileResponse {
	return &model.FileResponse{
		ID:           file.ID.Hex(),
		FileName:     file.FileName,
		OriginalName: file.OriginalName,
		FilePath:     file.FilePath,
		FileSize:     file.FileSize,
		FileType:     file.FileType,
		UploadedAt:   file.UploadedAt,
	}
}

// Inisialisasi service (dipanggil dari main.go atau routes)
var fileSvc FileService

func InitFileService() {
	fileSvc = NewFileService(repository.NewFileRepository(database.MongoDB), "./uploads")
}

// Wrapper untuk route
func UploadFile(c *fiber.Ctx) error    { return fileSvc.UploadFile(c) }
func GetAllFiles(c *fiber.Ctx) error   { return fileSvc.GetAllFiles(c) }
func GetFileByID(c *fiber.Ctx) error   { return fileSvc.GetFileByID(c) }
func DeleteFile(c *fiber.Ctx) error    { return fileSvc.DeleteFile(c) }