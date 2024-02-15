package domain

import "github.com/JMjirapat/qrthrough-api/internal/core/dto"

type LiffService interface {
	SignUp(body dto.RegisterRequestBody, lineID string) error
	GetAlumni(id int) (*dto.AlumniResponseBody, error)
}
