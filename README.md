# `gorun` - run all the GO things at once

Do you have a shitload of **GO** microservices all part of 1 project?  
Well in that case this program can solve the problem of 20 open terminal tabs. 

### Install
```sh
go install github.com/mjarkk/gorun@master
```

### Create a config
Create a file named `.gorun` or in the parrent dir create a file with the name of the folder with a `.` before it,  
Add a config like this:
```json
{
  "": {
    "server": ".",
    "email": "workers/email/Email.go",
    "notifi": "workers/notifications/Notifications.go",
    "paymnt": "workers/payments/Payments.go",
    "export": "workers/exports/Exports.go"
  }
}
```
Now when you execute `gorun`, it will run the programs spesified above in paralel and add a prefix to the console output of every program.  

### Custom config
You can also add custom configs like this:
```json
{
  "minimal": {
    "server": ".",
    "notifi": "workers/notifications/Notifications.go",
  }
}
```
And use the config via `gorun minimal`

### Q and A:

**Q: Can i add launch arguments?**  
A: Yes, although escape characters and qoutes do not work at the moment  

**Q: Can i add shell variables?**  
A: Yes just prefix the launch command with something like `SOME_SHELL_VAR=some_value` and it will work,  
But note, qoutes, escape characters and shell variables inside the shell variable will not work.  

**Q: Program X it's output is now broken, will you fix it?**  
A: TL;DR: No,  
Buttt, if it fixes the output for a lot of things maybe.  
The main purpose of this program is to run a lot of little microservices that just log garbage and to know from which micro service came the log.  
This is not made for programs that use special printing tricks
