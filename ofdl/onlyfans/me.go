package onlyfans

func (o *OnlyFans) Following() (*Following, error) {
	out := Following{}
	r, err := o.New().Get("lists/following").Request()
	if err != nil {
		return nil, err
	}

	if err := o.sign(r); err != nil {
		return nil, err
	}

	err = o.do(r, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type Following struct {
	UsersCount int `json:"usersCount"`
}
