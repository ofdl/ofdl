package onlyfans

func (o *OnlyFans) Following() (*Following, error) {
	out := Following{}
	outErr := map[string]interface{}{}
	r, err := o.New().Get("lists/following").Request()
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

	return &out, nil
}

type Following struct {
	UsersCount int `json:"usersCount"`
}
