# URL-Shortner
URL-Shortner in Go

# Setup Steps
1. Clone the repo
2. Set the environment variable for mongodb_url

    `export MONGO_URL="mongodb_url"`
  
3. Create a mongodb database by name **url_shortner**
4. Build go binary and run
5. Access the endpoint http://localhost:8000/encode and add payload:
     ``` 
     {
    "url": "http://<url>"
    }
    ```
    
6. To decode the url hit the endpoint http://localhost:8000/decode and add payload
    ```
    {
    "url":"<short_url>"
    }
    ```
    
