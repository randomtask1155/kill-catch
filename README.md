

## Purpose

simple app that listens on $PORT and will sleep for $KILL_SLEEP seconds before exiting when a SIGQUIT, SIGTERM, or SIGKILL is received.

## push app

```
cf push
```

## change how long the app sleeps

example making the app sleep for 60 seconds.

```
cf set-env kill-catch KILL_SLEEP 60
```

## kill app

stop the app and observe it hang when SIGTERM is received.

```
cf stop kill-catch
```

```
   2020-02-19T17:01:05.49-0600 [API/0] OUT Updated app with guid 2b407ee0-539c-4439-908d-d234012daee1 ({"state"=>"STOPPED"})
   2020-02-19T17:01:05.48-0600 [CELL/0] OUT Cell eebffebb-e67a-4a4a-bd8e-9475241a7de9 stopping instance 659eef47-87be-4e01-5c1c-2705
   2020-02-19T17:01:10.81-0600 [APP/PROC/WEB/0] ERR 2020/02/19 23:01:10 Got signal: terminated
   2020-02-19T17:01:10.81-0600 [APP/PROC/WEB/0] ERR 2020/02/19 23:01:10 Sleeping for 5s
   2020-02-19T17:01:10.81-0600 [CELL/SSHD/0] OUT Exit status 0
   2020-02-19T17:01:15.81-0600 [APP/PROC/WEB/0] ERR 2020/02/19 23:01:15 Sleep Complete
   2020-02-19T17:01:15.81-0600 [APP/PROC/WEB/0] OUT Exit status 0
   2020-02-19T17:01:15.82-0600 [CELL/0] OUT Cell eebffebb-e67a-4a4a-bd8e-9475241a7de9 destroying container for instance 659eef47-87be-4e01-5c1c-2705
   2020-02-19T17:01:16.00-0600 [PROXY/0] OUT Exit status 137
   2020-02-19T17:01:16.10-0600 [CELL/0] OUT Cell eebffebb-e67a-4a4a-bd8e-9475241a7de9 successfully destroyed container for instance 659eef47-87be-4e01-5c1c-2705
```