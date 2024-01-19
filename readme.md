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