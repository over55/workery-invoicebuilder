package dtos

type GreetingRequestDTO struct {
	Name string `json:"name"`
}

type GreetingResponseDTO struct {
	Message string `json:"message"`
}
