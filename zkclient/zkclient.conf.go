package main

const (
	defaultConfigFileContent = `
# root directory for zkclient settings, logs etc
root = ~/.zkclient

# print certificate fingerprint
tlsverbose = yes

# annoy user by beeping on incoming messages
# beep = yes

# Draw separator to show where conversation left off
# separator = yes

# logging and debug
[log]

# savehistory saves commands but not text to a history file
savehistory = no

# timeformat for logging purposes
# see https://golang.org/pkg/time/#Time.Format for more details
timeformat = 15:04:05

# longtimeformat is used when a year/month/day is required.
longtimeformat = 2006-01-02 15:04:05

# logfile contains log file name location
logfile = ~/.zkclient/zkclient.log

# enable/disable debug output to log
debug = no

# launch go's profiler on specified url
# requires debug = yes
profiler = 127.0.0.1:6061

# Valid ui colors: black, red, green, yellow, blue, magenta, cyan and white
# Valid atttributes are: none, underline and bold
# format is: attribute:foreground:background
[ui]
nickcolor = bold:white:black
gcothercolor = bold:green:black
pmothercolor = bold:cyan:black
`
)
