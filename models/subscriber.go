package models

// SubscriberResponse : Structure format in db
type SubscriberResponse struct {
	ID   int    `json:"id"`
	Mail string `json:"mail"`
	Name string `json:"name"`
}

// SubscriberFindBy : Get subscriber according to params
func SubscriberFindBy(key string, val string) (SubscriberResponse, error) {
	var sub SubscriberResponse

	err := DB.QueryRow("SELECT id, mail, name FROM subscriber WHERE "+key+" = ?", val).Scan(&sub.ID, &sub.Mail, &sub.Name)

	if err != nil {
		return sub, err
	}

	return sub, nil
}
