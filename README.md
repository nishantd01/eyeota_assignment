# Running Locally.

Using tarball
1. If not done , Uncompress tar.gz eyeota_assignment.tar.gz into go workspace 
2. mkdir nishantd01 inside $HOME/go/src/github.com , so the directory looks like $HOME/go/src/github.com/nishantd01
3. Move eyeota_asignment folder inside nishantd01
4. Go to eyeota_assignment folder
5. Run cmd **dep init**
6. Run cmd **dep ensure**
7. Run cmd **go build**
8. Run binary using cmd **./eyeota_assignment**
9. For running test cases run cmd **go test**


Alternatively Using github repo
1. You can directly get the changes commited from github. by running cmd `go get github.com/nishantd01/eyeota-assignment`
2. Go inside folder src/nishantd01/eyota-assignment
3. Run cmd **dep init**
4. Run cmd **dep ensure**
5. Run cmd **go build**
6. Run binary using cmd **./eyeota_assignment**
7. For running test cases run cmd **go test**



**NOTE** : I haven't used any third party libraries or packages.


# Data Structure Used
3D Map Named _DataMap_ with orgName,ParamName & ParamVal as keys and value is array of SegmentConfig containing Ids.

# Thought Process
My thought process was to initialize the data.json into a **_3D map_**, so that look up will be very fast when queried for values based on keys.

Alternatives I considered was ->
- to keep data in **_redis_** so that again & again populate was not required once done unless or untill redis crashes
- to keep data in **_database such as Mysql Or Postgres_** , but query could have been slow as creating connection & fetching data would have taken some fraction of seconds depending upon the machine.


So I considered **_3D map_**  as best approach since data is already there in data.json so again don't need to save it in db or populate it in redis cache, directly populating it a map would do the work & look up will be fast.


# Code Walk through
The folders data & lookupcache kept in the sample project are copied into eyeota_assignment so that interface can be implemented & data.json can be read.

1. main.go  -> It reads the data.json in function _initializeCache(fileName)_ which returns a struct containing dataMap(orgName,paramName, paramVal as Key & arrayOf Segment Config containing Ids) which can later be used to retrive info. 
It also exposes an end point **/api/v1/getSegmentByQuery** so that the request can be tested using postman. function _getSegment_ is a handler function to entertain this GET request.

2. struct.go -> It contains all the generic struct , Request Params , Response Params & DataMap 
3. fetchCache.go -> This file implements GetSegmentForOrgAndKey & GetSegmentForOrgAndKeyAndVal  interface function and sends back the result(array of segmentConfig) to the callee.
4. fetchCache_test.go -> This file contains unit tests for testData.json file kept inside data folder. It invokes GetSegmentForOrgAndKey & GetSegmentForOrgAndKeyAndVal methods and checks if the output is correct as per testData.json or not , if anything fails it reports with failure message.




    





# Testing

Run the binary **_eyeota_assignment_** and when binary is running it runs on 7020 port.
So the query can be tested using postman or insomnia client

API URL is **localhost:7020/api/v1/getSegmentByQuery**

**_Method_** : GET
Sample 1
**_Request Params_**

```
{
    "orgName":"6lkb2cv",
    "paramName":"testedu",
    "paramVal":""
}
```

**_Response Params_**

```
[{"Id":"n277"}]
```

Sample 2

```
{
    "orgName":"6lkb2cv",
    "paramName":"",
    "paramVal":""
}
```

**_Response Params_**

```
Invalid Input
```