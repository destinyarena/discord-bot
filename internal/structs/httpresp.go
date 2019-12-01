package structs

type (
    ResponseBody struct {
        Payload ResponseBodyPayload `json:"payload"`
    }

    ResponseBodyPayload struct {
        Code string `json:"code"`
    }
)
