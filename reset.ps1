# just run this once
$executablePath = ".\startallback_trialreset.exe"
$taskName = "StartAllBack trial reset"
$trigger = New-ScheduledTaskTrigger -Weekly -DaysOfWeek Monday -WeeksInterval 1
$action = New-ScheduledTaskAction -Execute $executablePath
Register-ScheduledTask -Action $action -Trigger $trigger -TaskName $taskName -User "NT AUTHORITY\SYSTEM" -Force