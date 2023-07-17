package onlyfans

import "fmt"

func (o *OnlyFans) GetMessages(uid int, id *int) (*PagableList[Message], error) {
	page := &PagableList[Message]{}
	r, err := o.New().
		Get(fmt.Sprintf("chats/%d/messages", uid)).
		QueryStruct(&GetMessagesOptions{
			Limit:     50,
			Order:     "desc",
			SkipUsers: "all",
			ID:        id,
		}).
		Request()
	if err != nil {
		return nil, err
	}

	if err := o.sign(r); err != nil {
		return nil, err
	}

	_, err = o.Do(r, page, nil)
	if err != nil {
		return nil, err
	}

	return page, nil
}

type GetMessagesOptions struct {
	Limit     int    `url:"limit,omitempty"`
	Order     string `url:"order,omitempty"`
	SkipUsers string `url:"skip_users,omitempty"`
	ID        *int   `url:"id,omitempty"`
}

type Message struct {
	ID       int `json:"id"`
	FromUser struct {
		ID int `json:"id"`
	} `json:"fromUser"`
	CreatedAt  string         `json:"createdAt"`
	MediaCount int            `json:"mediaCount"`
	Media      []MessageMedia `json:"media"`
	Text       string         `json:"text"`
}

type MessageMedia struct {
	ID      int     `json:"id"`
	Type    string  `json:"type"`
	CanView bool    `json:"canView"`
	Src     *string `json:"src"`
}

type PagableList[T any] struct {
	HasMore bool `json:"hasMore"`
	List    []T  `json:"list"`
}
