# Caster
A RestAPI for sending emails.

```
   _________   _____________________ 
  / ____/   | / ___/_  __/ ____/ __ \
 / /   / /| | \__ \ / / / __/ / /_/ /
/ /___/ ___ |___/ // / / /___/ _, _/ 
\____/_/  |_/____//_/ /_____/_/ |_|  
```

## Installation
1. clone the repo
```
git clone https://github.com/TRudenko22/caster
cd caster
```

2. build the docker image
```
docker build -t caster .
```

3. run the docker image
```
docker run -p 8080:9000 -p 587:587 -e EMAIL_PASSWORD=<your gmail app password> caster
```

## Usage

The docker image has the two default environment variables set:
```
EMAIL_SERVER  = smtp.gmail.com
EMAIL_PORT    = 587
```

These can be changed by passing them in the docker run command and publishing their respective ports

```
docker run -p 8080:9000 -p <EMAIL_PORT>:<EMAIL_PORT> -e EMAIL_SERVER=<your email server> -e EMAIL_PORT=<your email port> -e EMAIL_PASSWORD=<your gmail app password> caster
```

Additionally you can use the provided `docker-compose.yml` as a template for your own deployment. 

The API has two endpoints:
`/`     - GET - returns a simple message for testing
`/send` - POST - sends an email

The `/send` endpoint requires a JSON body with the following fields:
```
{
    "sender": <your email address>,
    "recipients": [<list of email addresses>],
    "subject": <email subject>,
    "body": <email body>
}
```

## Roadmap
- [ ] more robust body parsing
- [ ] better error handling
- [ ] more email providers
- [ ] client application in Go

