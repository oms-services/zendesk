# _Zendesk_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/oms-services/zendesk.svg?branch=master)](https://travis-ci.com/oms-services/zendesk)
[![codecov](https://codecov.io/gh/oms-services/zendesk/branch/master/graph/badge.svg)](https://codecov.io/gh/oms-services/zendesk)

An OMG service for zendesk, it offers one of the most efficient and user-friendly Help Desk features available. It allows companies to design their own online Help Center format to ensure that it matches their unique brand message and customers' needs.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Create User
```coffee
zendesk createUser name:'name'
{"id": id,"url": "url for user","name": "user name","alias": "alias name","created_at": "2019-08-09T15:17:19Z","updated_at": "2019-08-09T15:17:19Z","active": true,"verified": false,  "shared": false,"shared_agent": false,"locale": "en-US","locale_id": 1,"time_zone": "Eastern Time (US & Canada)","email": "abc@example.com","phone": "+911234567890","details": "some details about user","notes": "Any notes you want to store about the user","role": "end-user","moderator": false,
"ticket_restriction": "requested","only_private_comments": false,"restricted_agent": true,"suspended": false}
```
##### Create Ticket
```coffee
zendesk createTicket subject:'subject' description:'description' requesterId:'requesterId'
{"id": id,"url": "ticket url","external_id": "","type": "task","subject": "OMG Ticket", "raw_subject": "OMG Ticket","description": "ticket from OMG","priority": "high","status": "new",  "recipient": "abc@example.com","requester_id": requester_id,"submitter_id": submitter_id,  "group_id": group_id,"has_incidents": false,"via": {"via object"},"created_at": "2019-08-09T15:20:35Z","updated_at": "2019-08-09T15:20:35Z","brand_id": brand_id}
```
##### List All Tickets
```coffee
zendesk listTicket  sortBy:'sortBy'
{"Comments": null,"Tickets": ["List of Tickets"],"Users": null,"Groups": null,"Audits": null,  "NextPage": null,"PreviousPage": null,"Count": Ticket Count}
```
##### Delete Ticket
```coffee
zendesk deleteTicket  ticketId:'ticketId'
{"success": true,"message": "Ticket deleted successfully","statusCode": 200}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create User
```shell
$ omg run createUser -a email=<EMAIL> -a name=<NAME> -a alias<USER_ALIAS> -a details=<DETAILS> -a notes=<NOTES> -a phone=<PHONE_NUMBER> -e DOMAIN_NAME=<DOMAIN_NAME> -e EMAIL=<EMAIL> -e API_TOKEN=<API_TOKEN>
```
##### Create Ticket
```shell
$ omg run createTicket -a type=<TYPE> -a subject=<SUBJECT> -a rawSubject=<RAW_SUBJECT> -a description=<DESCRIPTION> -a priority=<TICKET_PRIORITY> -a status=<TICKET_STATUS> -a recipient=<RECIPIENT_EMAIL> -a requesterId=<REQUESTER_ID> -e DOMAIN_NAME=<DOMAIN_NAME> -e EMAIL=<EMAIL> -e API_TOKEN=<API_TOKEN>
```
##### List All Tickets
```shell
$ omg run listTicket -a sortBy=<SORT_BY> -e DOMAIN_NAME=<DOMAIN_NAME> -e EMAIL=<EMAIL> -e API_TOKEN=<API_TOKEN>
```
##### Delete Ticket
```shell
$ omg run deleteTicket -a ticketId=<TICKET_ID> -e DOMAIN_NAME=<DOMAIN_NAME> -e EMAIL=<EMAIL> -e API_TOKEN=<API_TOKEN>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/zendesk/blob/master/LICENSE).
