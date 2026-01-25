package userhttp

import (
	"time"

	"github.com/PrinceM13/knowledge-hub-api/internal/user"
)

// list item response

type UserListItemResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func toUserListItem(u *user.User) UserListItemResponse {
	return UserListItemResponse{
		Email: u.Email,
		Name:  u.Name,
	}
}

func toUserListItems(users []*user.User) []UserListItemResponse {
	out := make([]UserListItemResponse, 0, len(users))

	for _, u := range users {
		out = append(out, toUserListItem(u))
	}

	return out
}

// detail response

type UserDetailResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func toUserDetail(u *user.User) UserDetailResponse {
	return UserDetailResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
}
