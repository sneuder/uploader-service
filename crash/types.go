package crash

const (
	NotParamsProvided   = 10
	AuthHeaderNotValid  = 20
	MissingSomeFields   = 50
	ValidationData      = 51
	DifferentFieldTypes = 52
	DifferentPassword   = 101
	NotFoundUserInAuth  = 102
	FailToUpadeUSer     = 103
	FailToGetFile       = 201
	FailToProcessFile   = 202
)

var ErrorCodes = map[int]string{
	NotParamsProvided:  "Not params provided",
	AuthHeaderNotValid: "Authorization Header not valid",
	MissingSomeFields:  "Missing some fields",
	ValidationData:     "Some fields are not completed",
	DifferentPassword:  "Different passwords",
	NotFoundUserInAuth: "User not found in auth",
	FailToUpadeUSer:    "Fail to update user",

	FailToProcessFile: "Fail to process file",
	FailToGetFile:     "Fail to get file",
}
