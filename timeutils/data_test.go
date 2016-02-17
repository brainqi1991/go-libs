package timeutils

import "testing"

func Test_GetTimestamp(t *testing.T) {
	t.Log(GetTimestamp())
}

func Test_GetTimestampString(t *testing.T) {
	t.Log(GetTimestampString())
}

func Test_GetTimestampInMilli(t *testing.T) {
	t.Log(GetTimestampInMicro())
}

func Test_GetTimestampInMilliString(t *testing.T) {
	t.Log(GetTimestampInMicroString())
}

func Test_GetTimestampInMicro(t *testing.T) {
	t.Log(GetTimestampInMicro())
}

func Test_GetTimestampInMicroString(t *testing.T) {
	t.Log(GetTimestampInMicroString())
}

func Test_GetCurrentTimeFormat(t *testing.T) {
	t.Log(GetCurrentTimeFormat(DATE_TIME_FMT))
	t.Log(GetCurrentTimeFormat(DATE_FMT))
	t.Log(GetCurrentTimeFormat(TIME_FMT))

	t.Log(GetCurrentTimeFormat(DATE_TIME_FMT_CN))
	t.Log(GetCurrentTimeFormat(DATE_FMT_CN))
	t.Log(GetCurrentTimeFormat(TIME_FMT_CN))
}

// function without any arguments or returned values
func play() {
	// empty function
}

func Test_Elapse(t *testing.T) {
	t.Logf("execute GetTimeStamp function nano time:%v", Elapse(play))
}
