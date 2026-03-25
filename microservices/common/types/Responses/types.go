package responses

type GenericResponse struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}
