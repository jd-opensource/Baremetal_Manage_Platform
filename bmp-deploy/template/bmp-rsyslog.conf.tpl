module(load="imudp") # needs to be done just once
input(type="imudp" port="514")

$template RemoteLogs,"/var/log/bmp/bmp-rsyslog/%fromhost-ip%/%PROGRAMNAME%-%$YEAR%-%$MONTH%-%$DAY%.log"
$template DynamicDir,"/var/log/bmp/bmp-rsyslog/%fromhost-ip%"
:syslogtag, startswith, "ip" ?DynamicDir
if $fromhost-ip != "127.0.0.1" then ?RemoteLogs
& ~