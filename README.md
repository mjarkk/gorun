# `gorun` - run all the things at once

Do you have a shitload of microservices all part of 1 project?  
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