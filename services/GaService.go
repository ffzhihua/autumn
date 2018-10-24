package services

import (
	"autumn/tools/ga"
)

type GaService struct {

}

func (i *GaService) FirstVerify(secret string, code string) bool  {

	gac := ga.Code(secret)

	if code == gac {
		return true
	}

	return false
}
