# WASAPhoto - Keep in touch with your friends by sharing photos of special moments.

## about
WASAPhoto is web application based on RESTfull architecture to allows upload your photos directly from your PC, and they will be visible to everyone following you.


## Roadmap

1. [X] define APIs using the [https://www.openapis.org/](OpenAPI) standard
2. [ ] design and develop the server side (“backend”) in [https://go.dev/](Go)
3. [ ] design and develop the client side (“frontend”) in [https://developer.mozilla.org/en-US/docs/Web/JavaScript?retiredLocale=it](JavaScript)
4. [ ] create a [https://www.docker.com/](Docker) container image for deployment

## Functional design specification to be implemented

1. WASAPhoto users actions
    - stream of photos --> add or remove (all information on photo) own photo uploaded
        - [X] API
        - [ ] implementation

    - [ ] followers --> list other user that follow a user
        - [X] API
        - [ ] implementation

    - [ ] following --> user can place or unplace follow to other users
        - [X] API
        - [ ] implementation

    - [ ] likes --> user can assign like to photos from others users 
        - [X] API
        - [ ] implementation

    - [ ] banned users --> user can banned or unbanned other users
        - [X] API
        - [ ] implementation

    - [ ] username --> user can modify his username
        - [X] API
        - [ ] implementation

    - [ ] comments --> user can comment or uncomment other user photo's 
        - [X] API
        - [ ] implementation

    - [ ] search --> searching other users via username 
        - [X] API
        - [ ] implementation

    - [ ] login (Simplified login) --> login on register on WASAPhoto with username 
        - [X] API
        - [ ] implementation

2. Photos datas
    - [ ] data upload
        - [X] API
        - [ ] implementation

    - [ ] time upload 
        - [X] API
        - [ ] implementation

    - [ ] likes recived
        - [X] API
        - [ ] implementation

    - [ ] comments recived
        - [X] API
        - [ ] implementation

3. System specifications

    - [ ] Sort photo stream in reverse order by data  and time update value
        - [X] API
        - [ ] implementation

