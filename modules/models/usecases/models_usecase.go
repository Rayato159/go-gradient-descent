package usecases

import "www.github.com/Rayato159/go-gradient-descent/modules/entities"

type modelsUse struct {
	ModelsRep entities.ModelsRepository
}

func NewModelsUsecase(modelsRep entities.ModelsRepository) *modelsUse {
	return &modelsUse{
		ModelsRep: modelsRep,
	}
}
