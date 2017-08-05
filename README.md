# Graylog Hook for [Logrus](https://github.com/sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" />&nbsp;

Use this hook to send your logs to [Graylog](http://graylog2.org) server over UDP.
The hook is non-blocking: even if UDP is used to send messages, the extra work
should not block the logging function.

All logrus fields will be sent as additional fields on Graylog.

This is forked from gemnasium/logrus-graylog-hook, on the top of it , I changed it to fulfill my need, and with a few optimizations
1. Gzip and zlib compression pool
2. bytes.Buffer pool 

## Usage

The hook must be configured with:

* A Graylog GELF UDP address (a "ip:port" string).
* an optional hash with extra global fields. These fields will be included in all messages sent to Graylog

### Disable standard logging

For some reason, you may want to disable logging on stdout, and keep only the messages in Graylog (ie: a webserver inside a docker container).
You can redirect `stdout` to `/dev/null`, or just not log anything by creating a `NullFormatter` implementing `logrus.Formatter` interface:
