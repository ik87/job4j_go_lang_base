package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)

	if req == nil {
		res = append(res, "req is nil")
		return res
	}

	if req.UserID == "" {
		res = append(res, "req.UserID is empty")
	}

	if req.Title == "" {
		res = append(res, "req.Title is empty")
	}

	if req.Description == "" {
		res = append(res, "req.Description is empty")
	}

	return res
}
