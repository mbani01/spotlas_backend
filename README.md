# Query
This query is used to return spots which have a domain with a count greater than 1. It also modifies the website field to only contain the domain and counts the number of spots with the same domain.
# Endpoint
This endpoint is built using Golang, Fiber, Postgres and Postgis. It takes four parameters: Latitude, Longitude, Radius (in meters) and Type (circle or square) and returns all spots within the specified area ordered by distance.

### Running the endpoint
Make sure you have Go, Postgres and Postgis installed on your machine.\
Clone this repository.\
Update the database credentials in the connection.go file to match your Postgres setup.\
Run go run main.go to start the endpoint.\
The endpoint will be listening on the specified port in the main.go file (default is 3000)\
You can now make requests to the endpoint as described in the Usage section above.

### Usage
To use this endpoint, make a GET request to the endpoint URL with the following parameters:

Latitude: The latitude of the center point of the area to search for spots.\
Longitude: The longitude of the center point of the area to search for spots.\
Radius: The radius of the area to search for spots (in meters).\
Type: The type of area to search for spots. This parameter should be set to "circle" or "square" .

The endpoint will return an array of objects containing all fields in the data set.

Example
To search for spots within a radius of 250 meters of the point (latitude: 51.509865, longitude: -0.118092), the following request can be made:

GET http://localhost:3000/?Latitude=51.509865&Longitude=-0.118092&radius=250&type=circle
