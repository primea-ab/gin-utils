package elk46

type Status string

const (
	StausCreated   Status = "created"   // We've received your request.
	StatusSent            = "send"      // We've sent your message to the recipient
	StatusFailed          = "failed"    // Unable to deliver your message
	StatuDelivered        = "delivered" // Message has reached recipient's phone.
)

type Response struct {
	Id            string `json:"id"`
	Status        Status `json:"status"`
	From          string `json:"from"`
	To            string `json:"to"`
	Parts         int    `json:"parts"`
	Message       string `json:"message"`
	Created       string `json:"created"`
	Delivered     string `json:"delivered"`
	Cost          int    `json:"cost"`
	EstimatedCost int    `json:"estimated_cost"`
	Direction     string `json:"direction"`
	DontLog       string `json:"dontLog"`
}

type DeliveryReport struct {
	// 46elk message id
	Id string `json:"id"`
	// Either ”delivered” or ”failed”.
	Status Status `json:"status"`
	// The delivery time in UTC. Only included if status is set to delivered.
	Delivered string `json:"delivered"`
}
