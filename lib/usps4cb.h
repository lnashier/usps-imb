/*******************************************************************************
** EncAPI.h
**
*******************************************************************************/
 
/*******************************************************************************
**
** Return Codes
**
*******************************************************************************/
 
#define USPS_FSB_ENCODER_API_SUCCESS                           0
#define USPS_FSB_ENCODER_API_SELFTEST_FAILED                   1
#define USPS_FSB_ENCODER_API_BAR_STRING_IS_NULL                2
#define USPS_FSB_ENCODER_API_BYTE_CONVERSION_FAILED            3
#define USPS_FSB_ENCODER_API_RETRIEVE_TABLE_FAILED             4
#define USPS_FSB_ENCODER_API_CODEWORD_CONVERSION_FAILED        5
#define USPS_FSB_ENCODER_API_CHARACTER_RANGE_ERROR             6
#define USPS_FSB_ENCODER_API_TRACK_STRING_IS_NULL              7
#define USPS_FSB_ENCODER_API_ROUTE_STRING_IS_NULL              8
#define USPS_FSB_ENCODER_API_TRACK_STRING_BAD_LENGTH           9
#define USPS_FSB_ENCODER_API_TRACK_STRING_HAS_INVALID_DATA    10
#define USPS_FSB_ENCODER_API_TRACK_STRING_HAS_INVALID_DIGIT2  11
#define USPS_FSB_ENCODER_API_ROUTE_STRING_BAD_LENGTH          12
#define USPS_FSB_ENCODER_API_ROUTE_STRING_HAS_INVALID_DATA    13
 
int USPS4CB(char *TrackPtr, char *RoutePtr, char *BarPtr);
int USPSVCB(char *TrackPtr, char *RoutePtr, char *BarPtr);
