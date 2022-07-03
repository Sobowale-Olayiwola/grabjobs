# GrabJobs Technical Assignment
Assignment Type: Backend API

Provided Data List: The assignment includes 2000 job titles and coordinates from Singapore

Tasks to be achieved:
- Implement nearby 1km - 2km - 5km job location title request location 
- Search feature with job title

## Technology:
You may use your preferred programming language but we recommended with GoLang. Please take note that you will need to provide us a working and testable website link and source code after completion of this assessment.

For example API: 
http://abc.com/near_by?lat=1.290270&lng=103.851959&radius=1
Expected result :  list of data return with json format
 
http://abc.com/near_by?title=Accountant&lat=1.290270&lng=103.851959&radius=1 
Expected result :  list of data return with json format

## Actual Implementations
- Baseurl https://ola-grabjobs.herokuapp.com

### Endpoint
- https://ola-grabjobs.herokuapp.com/api/v1/jobs-locations/near-by?lat=1.290270&lng=103.851959&radius=1&title=accountant

- Response
```json
{
  "payload": [
    {
      "title": "ACCOUNTANT ASSISTANT",
      "Location": {
        "type": "Point",
        "coordinates": [
          103.847,
          1.2873
        ]
      }
    },
    {
      "title": "Financial Accountant – Iron Ore with a global MNC (7 to 11 yrs required)",
      "Location": {
        "type": "Point",
        "coordinates": [
          103.852,
          1.2832
        ]
      }
    }
  ]
}
```
- https://ola-grabjobs.herokuapp.com/api/v1/jobs-locations/near-by?lat=1.290270&lng=103.851959&radius=1

- Sample Response
```json
{
  "payload": [
    {
      "title": "Operation Assistant [Fish farm / 5.5 days / CCK] 9157",
      "Location": {
        "type": "Point",
        "coordinates": [
          103.852,
          1.29027
        ]
      }
    },
    {
      "title": "Welder cum Worker",
      "Location": {
        "type": "Point",
        "coordinates": [
          103.852,
          1.29027
        ]
      }
    }
  ]
}
```

## Test
The test covers the service layer and the helper function to load the csv file
go test -cover ./...

## To run the application
- cd cmd/api
- go build
- ./api
