package usps

// Encoder API status codes
const (
	StatusEncoderAPISuccess                    = 0
	StatusEncoderAPISelfTestFailed             = 1
	StatusEncoderAPIBarStringIsNull            = 2
	StatusEncoderAPIByteConversionFailed       = 3
	StatusEncoderAPIRetrieveTableFailed        = 4
	StatusEncoderAPICodewordConversionFailed   = 5
	StatusEncoderAPICharacterRangeError        = 6
	StatusEncoderAPITrackStringIsNull          = 7
	StatusEncoderAPIRouteStringIsNull          = 8
	StatusEncoderAPITrackStringBadLength       = 9
	StatusEncoderAPITrackStringHasInvalidData  = 10
	StatusEncoderAPITrackStringHasInvalidDigit = 11
	StatusEncoderAPIRouteStringBadLength       = 12
	StatusEncoderAPIRouteStringHasInvalidData  = 13
)

// StatusText returns a text for the Encoder API status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	switch code {
	case StatusEncoderAPISuccess:
		return "Success"
	case StatusEncoderAPISelfTestFailed:
		return "Self Test Failed"
	case StatusEncoderAPIBarStringIsNull:
		return "Bar String Is Null"
	case StatusEncoderAPIByteConversionFailed:
		return "Byte Conversion Failed"
	case StatusEncoderAPIRetrieveTableFailed:
		return "Retrieve Table Failed"
	case StatusEncoderAPICodewordConversionFailed:
		return "Codeword Conversion Failed"
	case StatusEncoderAPICharacterRangeError:
		return "Character Range Error"
	case StatusEncoderAPITrackStringIsNull:
		return "Track String Is Null"
	case StatusEncoderAPIRouteStringIsNull:
		return "Route String Is Null"
	case StatusEncoderAPITrackStringBadLength:
		return "Track String Bad Length"
	case StatusEncoderAPITrackStringHasInvalidData:
		return "Track String Has Invalid Data"
	case StatusEncoderAPITrackStringHasInvalidDigit:
		return "Track String Has Invalid Digit"
	case StatusEncoderAPIRouteStringBadLength:
		return "Route String Bad Length"
	case StatusEncoderAPIRouteStringHasInvalidData:
		return "Route String Has Invalid Data"
	default:
		return ""
	}
}
