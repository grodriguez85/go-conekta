package conekta

//ShippingLine shipping line struct
type ShippingLine struct {
	ID             string   `json:"id,omitempty"`
	Object         string   `json:"object,omitempty"`
	Amount         *int64   `json:"amount,omitempty"`
	TrackingNumber string   `json:"tracking_number,omitempty"`
	Carrier        string   `json:"carrier,omitempty"`
	Method         string   `json:"method,omitempty"`
	ParentID       string   `json:"parent_id,omitempty"`
	Metadata       Metadata `json:"metadata,omitempty"`
}

//ShippingLines shipping lines struct
type ShippingLines struct {
	Object  string         `json:"object,omitempty"`
	HasMore *bool          `json:"has_more,omitempty"`
	Total   int64          `json:"total,omitempty"`
	Data    []ShippingLine `json:"data,omitempty"`
}
