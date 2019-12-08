package structs

type ReqBody struct {
    EntityID string `json:"entity_id"`
    EntityType string `json:"entity_type"`
    Type string `json:"type"`
    MaxAge int `json:"max_age"`
    MaxUses int `json:"max_uses"`
}

type RolesPayload struct {
    Discord string `json:"discord"`
    Skillvl int `json:"skillvl"`
}
