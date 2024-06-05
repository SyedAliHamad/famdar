**Get All Users**

**Endpoint**: GET /users

**Description**: Retrieves a list of all users.

**Response**:

- Status Code: 200 OK
- Body: An array of user objects.

**Get User Stats**

**Endpoint**: GET /users/stats

**Description**: Retrieves statistics about users.

**Response**:

- Status Code: 200 OK
- Body: Statistics data about users.

**Get Nearby Users**

**Endpoint**: GET /users/nearby?lat={lat}&lng={lng}&radius={radius}

**Description**: Fetches users within a specified radius from a given location.

**Query Parameters**:

- lat (float, required): Latitude of the reference location.
- lng (float, required): Longitude of the reference location.
- radius (float, required): Radius in meters.

**Response**:

- Status Code: 200 OK
- Body: An array of user objects within the specified radius from the given location.

**Delete a User**

**Endpoint**: DELETE /users/{id}

**Description**: Deletes a user with the specified ID.

**Path Parameters**:

- id (integer, required): The ID of the user to delete.

**Response**:

- Status Code: 204 No Content