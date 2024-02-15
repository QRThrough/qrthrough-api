package dto

type RegisterRequestBody struct {
	StudentCode string `json:"student_code"`
	IDToken     string `json:"id_token"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Tel         string `json:"tel"`
}

type AlumniResponseBody struct {
	InAlumni    bool    `json:"in_alumni"`
	StudentCode string  `json:"student_code"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	Tel         *string `json:"tel"`
}

type ErrorMessageReq struct {
	Detail  string `json:"detail"`
	Message string `json:"message"`
}

type ErrorMessageRes struct {
	Detail  []string `json:"detail"`
	Message string   `json:"message"`
}

type OTPReq struct {
	Status string            `json:"status"`
	Token  string            `json:"token"`
	Refno  string            `json:"refno"`
	Code   int               `json:"code"`
	Errors []ErrorMessageReq `json:"errors"`
}

type OTPVerify struct {
	Status string            `json:"status"`
	Msg    string            `json:"message"`
	Code   int               `json:"code"`
	Errors []ErrorMessageRes `json:"errors"`
}

type GetOTPRequestBody struct {
	Tel string `json:"tel"`
}

type GetOTPResponseBody struct {
	Token string `json:"token"`
	Refno string `json:"refno"`
}

type VerifyOTPRequestBody struct {
	Token string `json:"token"`
	Pin   string `json:"pin"`
}
