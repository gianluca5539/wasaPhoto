### WasaPhoto by Gianluca Gabriel Monte ###

This is my project for the Web and Software Architecture course.
The website is (to a certain point) responsive and looks nice on almost all platforms.
With some small adjustments it could also work nice on a smartphone!


# Docker build instructions

To build the front-end:
docker buildx build -f Dockerfile.frontend -t wasa-frontend ./ 

To build the back-end:
docker buildx build -f Dockerfile.backend -t wasa-backend ./ 

# Docker run instructions

To run the front-end:
docker run -it --rm -p 80:80 wasa-frontend

To run the back-end:
docker run -it --rm -p 3000:3000 wasa-backend 


# How stuff works
### Login
During login an account may be created if it doesn't exist.
Anyway, information about the user is fetched and saved in localStorage.

### Stream
The homepage contains the stream, which is a collection of content from users followed by the user.
The user can go to his profile or logout, and look for a specific user by username if needed.

### Profile Page
In the profile page the user can check their stats and their posts. They can also post a new picture.
A user is able to change username, feeling, bio.
On another user's profile page the user can follow and/or ban them.

# Inner workings
### Follow user
Adds/removes an entry from a "service" table "follow" in the database.
This table is queried together with user information to retrieve the follow/followers list.

### Ban user
Works in the same way as the "follow" table. If a user is banned, the follow is removed.
This means that a user will not see any pictures from the user who banned them.
If a user looks for the user who banned them, it will tell "user not found".
It is still vulnerable to IDOR, but since the names of the pictures are the millis from 1970 it's near
impossible to get a picture via IDOR.

### Remove post
Remove post deletes also the comments and likes linked to that post, respectively in the "comment"
and "like" tables.

### New post
An entry in the "post" table is created, and the picture is converted from b64 to png. The image is
saved in data/pictures.


# Gimmicks

### Responsive website
The website should be fully responsive and should work on most platforms (even cellphones!)

