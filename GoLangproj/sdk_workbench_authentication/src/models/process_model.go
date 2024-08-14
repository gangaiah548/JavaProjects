package models

type ProcessDeploymentModel struct {
	ID           string `json:"_id,omitempty" validate:"-"`
	Key          string `json:"_key,omitempty" validate:"required"`
	Collection   string `json:"collection,omitempty" validate:"-"`
	Rev          string `json:"_rev,omitempty" validate:"-"`
	OldRev       string `json:"_oldRev,omitempty" validate:"-"`
	Title        string `json:"title,omitempty" validate:"required"`
	Description  string `json:"description,omitempty" validate:"-"`
	Definition   string `json:"definition,omitempty" validate:"required"`
	Status       string `json:"status,omitempty" validate:"required"`
	InstancePool int    `json:"instancePool,omitempty" validate:"required"`
}
