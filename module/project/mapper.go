package project

func ConvertProjectToDTO(projects []ProjectWithCreateBy) []*ProjectInfo {
	var projectsDTO []*ProjectInfo

	for _, project := range projects {
		projectsDTO = append(projectsDTO, &ProjectInfo{
			ProjectWithCreateBy: &ProjectWithCreateBy{
				ProjectPure: ConvertProjectToPureProject(&project),
				IsStar:      project.IsStar,
				CreateBy:    project.CreateBy,
			},
			// todo hardcode
			ProjectPermission: ProjectPermission{
				SourcePermission:   3,
				ViewPermission:     3,
				WidgetPermission:   3,
				VizPermission:      3,
				SchedulePermission: 3,
				SharePermission:    true,
				DownloadPermission: true,
			},
		})
	}
	return projectsDTO
}

func ConvertProjectToPureProject(project *ProjectWithCreateBy) *ProjectPure {
	return &ProjectPure{
		Id:           project.Id,
		Name:         project.Name,
		Description:  project.Description,
		Pic:          project.Pic,
		OrgId:        project.OrgId,
		UserId:       project.UserId,
		Visibility:   project.Visibility,
		StarNum:      project.StarNum,
		IsTransfer:   project.IsTransfer,
		InitialOrgId: project.InitialOrgId,
	}
}

func ConvertProjectDetailToProjectInfo(info *ProjectInfo, detail *ProjectDetail) {
	info.ProjectWithCreateBy = &ProjectWithCreateBy{
		ProjectPure: &ProjectPure{
			Id:           detail.Id,
			Name:         detail.Name,
			Description:  detail.Description,
			Pic:          detail.Pic,
			OrgId:        detail.OrgId,
			UserId:       detail.UserId,
			Visibility:   detail.Visibility,
			StarNum:      detail.StarNum,
			IsTransfer:   detail.IsTransfer,
			InitialOrgId: detail.InitialOrgId,
		},
		CreateBy: detail.CreateBy,
	}

	info.ProjectPermission = ProjectPermission{
		SourcePermission:   3,
		ViewPermission:     3,
		WidgetPermission:   3,
		VizPermission:      3,
		SchedulePermission: 3,
		SharePermission:    true,
		DownloadPermission: true,
	}

}
