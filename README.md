### Slackbot

The Slackbot is a service written in Go and leverages the power of the [slackapi](https://github.com/cixtor/slackapi) library to connect to Slack's RTM _(Real Time Messaging)_ websocket and react to events triggered by users or bots in the allowed channels.

You can find a list of implemented commands below and you can use them as a template to create more commands to match your specific needs. The project also includes a _Makefile_ with an instruction to facilitate the deployment of your code to production using the cross-compilation feature available in the Go compiler, rsync and a cronjob in the remote machine.

### Deployment

The _Makefile_ in the repository includes an instruction to deploy your new code to production, it will compile the source files in the host cross-compiling to a Linux-adm64 machine, then will upload the _"init.sh"_ to the remote server as a script into the `/etc/init.d/` directory, then will upload the new binary + README + auto-deploy script to a pre-defined directory. The target directory of the upload will be monitored by another script _"autodeploy.sh"_ which runs with a cronjob every minute and detects if the new binary, named _"slackbot.new"_ exists, in which case will assume that a new deployment was executed, then will stop the running service using the script in `/etc/init.d/`, replace the old binary with the new one, and start the service again using the new code.

Please make the appropriate changes to match your needs.

### Command - Help

Send `help` as a direct message to **@slackbot** and it will reply with the content of this Markdown file with the corresponding modifications to make the titles, links, bold text appear with the correct style in Slack.

### Command - Uptime

Send `uptime` as a direct message to **@slackbot** and it will reply with the amount of time since the last update, this is, the time since the service was restarted. The time is reset every time the cronjob in the remote machine detects a deployment.
