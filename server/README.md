# Assets/Liabilities Manager

This project is the server component to a generic assets/liabilities manager.
Users can create an account (Not verifying emails or providing functionality to reset passwords or anything fancy in this project)
Financial records can be created as either an asset or a liability, and the users total net worth will be calculated accordingly

## Requirements
- go1.13.x (Will most likely work with go1.11 or 1.12 (Needs module support), but testing has not been done with older versions)
- docker-compose

## Setup
- Create a file called .env with the following contents
```
AL_DB_HOST=localhost
AL_DB_PORT=5432
AL_DB_NAME=postgres
AL_DB_USER=postgres
AL_DB_PASSWORD=postgres
AL_DB_SSL_MODE=disable
AL_DEBUG_MODE=1
```

## Starting the server
- Start the database using ```docker-compose up -d```
- Run the migrations using ```make db-migrate```
- Start the server using ```make run```
- Server will listen on ```localhost:8080``` by default

## Routes

### Login
- URL: ```/auth/login```
- Description: Creates a user session
- Method: POST
- Body Parameters:
    ```
    {
        username: string[255]
        password: string[255]
    }
    ```
- Response:
    ```
    {
        id: string
        username: string
        fullname: string
        created_at: timestamp
        updated_at: timestamp
    }
    ```
- Cookies: A cookie with the name ```user``` will be created if the provided credentials were valid

### Logout
- URL: ```/auth/logout```
- Description: Destroys the current user's session
- Method: POST
- URL Parameters: None
- Response:
    ```
    {
        message: "Logged out"
    }
    ```
- Cookies: The ```user``` cookie's MAX_AGE will be set to -1 if it exists

### List all Records
- URL: ```/finances/records```
- Method: GET
- URL Parameters:
  - offset=[int]
  - limit=[int]
    - Max = 500
    - If provided limit is > 500, the limit will be set to 500
- Response:
    ```
    {
        records: [
            {
                id: string
                type: RECORD_TYPE
                name: string
                balance: float
                created_at: timestamp
                updated_at: timestamp
            }
        ],
        asset_total: int
        liability_total: int
        net_worth: int
    }
    ```

### Create New Record
- URL: ```/finances/records```
- Method: POST
- Body Parameters
    ```
    {
        type: RECORD_TYPE
        name: string[255]
        balance: float
    }
    ```
- Response:
    ```
    {
        id: string
    }
    ```


### Fetch record
- URL: ```/finances/records/:id```
- Method: GET
- URL Parameters: None
- Response:
    ```
    {
        id: string
        type: RECORD_TYPE
        name: string
        balance: float
        created_at: timestamp
        updated_at: timestamp
    }
    ```
- Description: Returns the financial record associated with the given id

### Update record
- URL: ```/finances/records/:id```
- Method: PUT
- Body Parameters:
    ```
    {
        id: string
        name: *string[255]
        balance: *float
    }
    ```
- Response:
    ```
    {
        id: string
        type: RECORD_TYPE
        name: string
        balance: float
        created_at: timestamp
        updated_at: timestamp
    }
    ```
- Description: Updates the financial record associated with the given id

### Delete record
- URL: ```/finances/records/:id```
- Method: DELETE
- Body Parameters:
    ```
    {
        id: string
    }
    ```
- Response:
    ```
    {}
    ```
- Description: Deletes the financial record associated with the given id