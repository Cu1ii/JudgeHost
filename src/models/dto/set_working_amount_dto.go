package dto

type SetWorkingAmountDTO struct {
	MaxWorkingAmount int  `form:"max_work_amount" json:"max_work_amount" validate:"required"`
	ForceSet         bool `form:"force_set" json:"force_set"`
}
