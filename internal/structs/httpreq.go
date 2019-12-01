package structs

type ReqBody struct {
    EntityID string `json:"entity_id"`
    EntityType string `json:"entity_type"`
    Type string `json:"type"`
    MaxAge int `json:"max_age"`
    MaxUsers int `json:"max_users"`
}
