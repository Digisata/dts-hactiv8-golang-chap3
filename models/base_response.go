package models

type SuccessResult struct {
	Code   int    `json:"code" example:"200"`
	Status string `json:"status" example:"Ok"`
	Data   any    `json:"data"`
}

type UnauthorizedResult struct {
	Code    int    `json:"code" example:"401"`
	Status  string `json:"status" example:"UNAUTHORIZED"`
	Message string `json:"message" example:"message"`
}

type BadRequestResult struct {
	Code    int    `json:"code" example:"400"`
	Status  string `json:"status" example:"BAD REQUEST"`
	Message string `json:"message" example:"message"`
}

type NotFoundResult struct {
	Code    int    `json:"code" example:"404"`
	Status  string `json:"status" example:"NOT FOUND"`
	Message string `json:"message" example:"message"`
}

type InternalServerErrorResult struct {
	Code    int    `json:"code" example:"500"`
	Status  string `json:"status" example:"INTERNAL SERVER ERROR"`
	Message string `json:"message" example:"message"`
}
