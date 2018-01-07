package huego

type Error struct {
	Type        string `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
