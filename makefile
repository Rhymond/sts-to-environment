build:
	go build -ldflags "-X main.RoleArn=<set-here> -X main.SerialNumber=<set-here> -X main.Profile=<set-here>"

install:
	go install -ldflags "-X main.RoleArn=<set-here> -X main.SerialNumber=<set-here> -X main.Profile=<set-here>"