# WASAPhoto - Keep in touch with your friends by sharing photos of special moments.

## about
WASAPhoto is web application based on RESTfull architecture to allows upload your photos directly from your PC, and they will be visible to everyone following you.


## Roadmap

1. [X] define APIs using the [https://www.openapis.org/](OpenAPI) standard
2. [X] design and develop the server side (“backend”) in [https://go.dev/](Go)
3. [ ] design and develop the client side (“frontend”) in [https://developer.mozilla.org/en-US/docs/Web/JavaScript?retiredLocale=it](JavaScript)
4. [ ] create a [https://www.docker.com/](Docker) container image for deployment

## Functional design specification to be implemented

1. WASAPhoto users actions
    - stream of photos --> add or remove (all information on photo) own photo uploaded
        - [X] API
        - [X] implementation

    - [X] followers --> list other user that follow a user
        - [X] API
        - [X] implementation

    - [X] following --> user can place or unplace follow to other users
        - [X] API
        - [X] implementation

    - [X] likes --> user can assign like to photos from others users 
        - [X] API
        - [X] implementation

    - [X] banned users --> user can banned or unbanned other users
        - [X] API
        - [X] implementation

    - [X] username --> user can modify his username
        - [X] API
        - [X] implementation

    - [X] comments --> user can comment or uncomment other user photo's 
        - [X] API
        - [X] implementation

    - [X] search --> searching other users via username 
        - [X] API
        - [X] implementation

    - [X] login (Simplified login) --> login on register on WASAPhoto with username 
        - [X] API
        - [X] implementation

2. Photos datas
    - [X] data upload
        - [X] API
        - [X] implementation

    - [X] time upload 
        - [X] API
        - [X] implementation

    - [X] likes recived
        - [X] API
        - [X] implementation

    - [X] comments recived
        - [X] API
        - [X] implementation

3. System specifications

    - [X] Sort photo stream in reverse order by data  and time update value
        - [X] API
        - [X] implementation

npm run build-prod && npm run preview
