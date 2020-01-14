package problem0359

type Logger struct {
	messages map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	messages := make(map[string]int)
	return Logger{messages}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if prevTimestamp, ok := this.messages[message]; ok {
		if timestamp-prevTimestamp >= 10 {
			this.messages[message] = timestamp
			return true
		} else {
			return false
		}
	}
	this.messages[message] = timestamp
	return true
}
