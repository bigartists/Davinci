package source

func ConvertSourcesToDTO(sources []Source) *[]SourceDTO {
	sourceDTOs := make([]SourceDTO, len(sources))
	for index, source := range sources {
		sourceDTO := SourceDTO{
			Id:          source.ID,
			Name:        source.Name,
			Description: source.Description,
			Type:        source.Type,
			ProjectId:   source.ProjectID,
		}
		//sourceDTOs = append(sourceDTOs, sourceDTO)
		sourceDTOs[index] = sourceDTO
	}

	return &sourceDTOs
}

func ConvertSourcesToSourceDetail(source *Source) *SourceDetail {
	sourceDetail := &SourceDetail{}
	sourceDetail.SourceDTO = &SourceDTO{
		Id:          source.ID,
		Name:        source.Name,
		Description: source.Description,
		Type:        source.Type,
		ProjectId:   source.ProjectID,
	}
	sourceDetail.Config = source.Config

	return sourceDetail
}

func ConvertTablesToMap(tables *[]string) []Item {
	// Map over the slice and add the "type" field
	items := make([]Item, len(*tables))
	for i, s := range *tables {
		items[i] = Item{
			Name: s,
			Type: "TABLE",
		}
	}
	return items
}
