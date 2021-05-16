package main

type Logger struct {
	logs map[string]int // map to map message->timestamp
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	// check if this is even in the map
	if t, ok := this.logs[message]; ok {
		// check if t was closer than 10 seconds ago
		if timestamp-t < 10 {
			// can't create
			return false
		} else {
			// can create
			this.logs[message] = timestamp
			return true
		}
	} else {
		// create it
		this.logs[message] = timestamp
		return true
	}
}
