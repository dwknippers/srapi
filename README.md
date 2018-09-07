# srapi
_Work in progress_
## Description
API for retrieving timetable information from Saxion University of Applied Sciences' timetable service using a custom webscraper.

---
## Interface
### Authentication
Please note that for every request, an access token must be provided to access timetables.

This access token can be found as a cookie called `saxion_roosters[access_token]` on the [timetable ervice website](https://roosters.saxion.nl) after successfully authention.

Please provide this token in the `Authorization` header like so:
```
Authorization: Bearer <Access Token>
```
Where _\<Access Token\>_ would be your access token.

---

### Retrieve timetable
localhost:8000/group/**{groupName}**/schedule/week/**{weekNumber}**

Where **groupName**, and **weekNumber** correspond to the group name and week number respectively.

Example output:
```
[
	{
 		"Date": "2018-09-10T00:00:00Z",
		"Events": [
			{
				"Start": "0000-01-01T08:30:00Z",
				"End": "0000-01-01T11:00:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>",
				"Type": "Practicum",
				"LocationCode": "<Location Code>"
			},
			{
				"Start": "0000-01-01T11:45:00Z",
				"End": "0000-01-01T13:15:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>",
				"Type": "Hoorcollege",
				"LocationCode": "<Location Code>"
			}
		]
	},
	{
		"Date": "2018-09-11T00:00:00Z",
		"Events": [
			{
				"Start": "0000-01-01T08:30:00Z",
				"End": "0000-01-01T11:00:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>",
				"Type": "Werkcollege",
				"LocationCode": "<Location Code>"
			},
			{
				"Start": "0000-01-01T11:45:00Z",
				"End": "0000-01-01T13:15:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>",
				"Type": "<Type Name>",
				"LocationCode": "<Location Code>"
			},
			{
				"Start": "0000-01-01T13:15:00Z",
				"End": "0000-01-01T14:45:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>",
				"Type": "Werkcollege",
				"LocationCode": "<Location Code>"
			}
		]
	},
	{
		"Date": "2018-09-12T00:00:00Z",
		"Events": []
	},
	{
		"Date": "2018-09-13T00:00:00Z",
		"Events": [
			{
				"Start": "0000-01-01T08:30:00Z",
				"End": "0000-01-01T11:00:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>",
				"Type": "Workshop",
				"LocationCode": "<Location Code>"
			}
		]
	},
	{
		"Date": "2018-09-14T00:00:00Z",
		"Events": [
			{
				"Start": "0000-01-01T08:30:00Z",
				"End": "0000-01-01T12:30:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>"
				"Type": "Werkcollege",
				"LocationCode": "<Location Code>"
			},
			{
				"Start": "0000-01-01T13:15:00Z",
				"End": "0000-01-01T14:45:00Z",
				"Course": "<Course Name>",
				"TeacherName": "<Teacher Name>",
				"Type": "<Type Name>",
				"LocationCode": "<Location Code>"
			}
		]
	}
]
```
