package dto

type SampleRequest struct {
	Field string `json:"field" validate:"notblank"`
}
