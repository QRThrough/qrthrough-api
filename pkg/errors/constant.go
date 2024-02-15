package errors

//line webhhook
const (
	LW_SYSTEM_ERROR_DESC       = "เกิดข้อผิดพลาด: ไม่สามารถสร้าง QR Code ได้"
	LW_CANT_CHECK_OC_TIME_DESC = "เกิดข้อผิดพลาด: ไม่สามารถสร้าง QR Code ได้ (CODE:LW00000)"
	LW_SYSTEM_OFF_DESC         = "แจ้งเตือน: ระบบอยู่ในช่วงปิดให้ใช้งาน"
	LW_PARSE_INT_FAILED_DESC   = "เกิดข้อผิดพลาด: กรุณาลองใหม่อีกครั้ง (CODE:LW00001)"
	LW_NOT_FOUND_USER_DESC     = "เกิดข้อผิดพลาด: ไม่พบข้อมูลผู้ใช้ในระบบ, กรุณาลงทะเบียนก่อน"
	LW_NOT_ACTIVE_USER_DESC    = "แจ้งเตือน: บัญชีของท่านยังไม่เปิดใช้งานหรือถูกปิดใช้งาน, กรุณาติดต่อเจ้าหน้าที่"
	LW_OUT_OF_TIME_DESC        = "แจ้งเตือน: อยู่นอกช่วงให้บริการ (%v น. - %v น.)"
)

// liff
const (
	LIFF_BAD_GATEWAY_CODE          = "LE00000"
	LIFF_NO_SUB_CODE               = "LE00001"
	LIFF_ALREADY_USED_ALUMNI_CODE  = "LE00002"
	LIFF_ATOI_FAILED_CODE          = "LE00003"
	LIFF_DUPLICATE_LINE_CODE       = "LE00004"
	LIFF_DUPLICATE_STUDENT_ID_CODE = "LE00005"
	LIFF_DUPLICATE_TEL_CODE        = "LE00006"
	LIFF_SIGN_UP_FAILED_CODE       = "LE00007"
)

// otp
const (
	TBS_BAD_GATEWAY_CODE  = "OE00000"
	TBS_MSISDN_ERROR_CODE = "OE00001"
	TBS_WRONG_OTP_CODE    = "OE00002"
	TBS_EXPIRED_OTP_CODE  = "OE00003"
)

// scanner
const (
	SC_SYSTEM_OFF_CODE       = "SC00000"
	SC_OUT_OF_TIME_CODE      = "SC00001"
	SC_QRCODE_NOT_FOUND_CODE = "SC00002"
	SC_QRCODE_EXPIRED_CODE   = "SC00003"
	SC_QRCODE_USED_UP_CODE   = "SC00004"
)

const (
	SYSTEM_ERROR_CODE = "SE00000"
)
