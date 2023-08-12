package onlyfans

func (o *OnlyFans) GetSubscriptions(limit, offset int) ([]Subscription, error) {
	out := []Subscription{}
	outErr := map[string]interface{}{}
	r, err := o.New().Get("subscriptions/subscribes").
		QueryStruct(GetSubscriptionsOptions{
			Limit:  limit,
			Offset: offset,
			Type:   "active",
			Sort:   "desc",
			Field:  "expire_date",
		}).Request()
	if err != nil {
		return nil, err
	}

	if err := o.sign(r); err != nil {
		return nil, err
	}

	_, err = o.Do(r, &out, &outErr)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type GetSubscriptionsOptions struct {
	Offset int    `url:"offset"`
	Limit  int    `url:"limit"`
	Type   string `url:"type"`
	Sort   string `url:"sort"`
	Field  string `url:"field"`
}

type Subscription struct {
	ID       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Header   string `json:"header"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
