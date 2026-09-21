package service

import (
	"fmt"
	"math"
	"strings"

	"github.com/google/uuid"
)

type GenPass interface {
	GeneratePassword(length int) (string, error)
}

type genPassService struct {
}

func NewGenPass() GenPass {
	return &genPassService{}
}

func (g *genPassService) GeneratePassword(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}
	dash := "-"
	result := strings.ReplaceAll(uuid.New().String(), dash, "")
	max_len := len(result)
	if length > max_len {
		var times int = int(math.Ceil(float64(length) / float64(max_len)))
		for i := 0; i < times-1; i++ {
			result += strings.ReplaceAll(uuid.New().String(), dash, "")
		}
	}
	result = result[:length]
	return result, nil
}
