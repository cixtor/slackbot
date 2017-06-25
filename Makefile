deploy:
	env GOOS=linux GOARCH=amd64 go build -o slackbot.new -- *.go

	rsync -av --rsh "ssh -i ~/.ssh/id_rsa" -- \
	init.sh "pi@192.168.1.55:/etc/init.d/slackbot"

	rsync -av --rsh "ssh -i ~/.ssh/id_rsa" -- \
	slackbot.new autodeploy.sh README.md \
	"pi@192.168.1.55:/srv/slackbot/"

	rm -- slackbot.new 2>/dev/null
