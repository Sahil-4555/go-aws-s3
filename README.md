# React Frontend with Go Backend Demo

This project is a small demo application that combines a frontend built with React and a backend built with Go. It provides functionalities to upload files, list objects, delete objects, generate signed URLs, and create buckets.

## Backend (Go)

The backend server is built with Go and provides the following API endpoints:

- `POST /upload-file`: Uploads media files.
- `GET /object-list`: Retrieves a list of objects.
- `DELETE /delete-object`: Deletes an object.
- `GET /get-signed-url`: Generates a signed URL.
- `POST /create-bucket/:bucket`: Creates a bucket.

## Frontend (React)

The frontend is built with React and interacts with the backend APIs to perform various operations.

