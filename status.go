package usps

const (
	StatusEncoderApiSuccess                    = 0
	StatusEncoderApiSelfTestFailed             = 1
	StatusEncoderApiBarStringIsNull            = 2
	StatusEncoderApiByteConversionFailed       = 3
	StatusEncoderApiRetrieveTableFailed        = 4
	StatusEncoderApiCodewordConversionFailed   = 5
	StatusEncoderApiCharacterRangeError        = 6
	StatusEncoderApiTrackStringIsNull          = 7
	StatusEncoderApiRouteStringIsNull          = 8
	StatusEncoderApiTrackStringBadLength       = 9
	StatusEncoderApiTrackStringHasInvalidData  = 10
	StatusEncoderApiTrackStringHasInvalidDigit = 11
	StatusEncoderApiRouteStringBadLength       = 12
	StatusEncoderApiRouteStringHasInvalidData  = 13
)

func StatusText(code int) string {
	switch code {
	case StatusEncoderApiSuccess:
		return "Success"
	case StatusEncoderApiSelfTestFailed:
		return "Self Test Failed"
	case StatusEncoderApiBarStringIsNull:
		return "Bar String Is Null"
	case StatusEncoderApiByteConversionFailed:
		return "Byte Conversion Failed"
	case StatusEncoderApiRetrieveTableFailed:
		return "Retrieve Table Failed"
	case StatusEncoderApiCodewordConversionFailed:
		return "Codeword Conversion Failed"
	case StatusEncoderApiCharacterRangeError:
		return "Character Range Error"
	case StatusEncoderApiTrackStringIsNull:
		return "Track String Is Null"
	case StatusEncoderApiRouteStringIsNull:
		return "Route String Is Null"
	case StatusEncoderApiTrackStringBadLength:
		return "Track String Bad Length"
	case StatusEncoderApiTrackStringHasInvalidData:
		return "Track String Has Invalid Data"
	case StatusEncoderApiTrackStringHasInvalidDigit:
		return "Track String Has Invalid Digit"
	case StatusEncoderApiRouteStringBadLength:
		return "Route String Bad Length"
	case StatusEncoderApiRouteStringHasInvalidData:
		return "Route String Has Invalid Data"
	default:
		return ""
	}
}
