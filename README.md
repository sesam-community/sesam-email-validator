# sesam-email-validator
Simple micro service to validate email addresses in Sesam.io powered applications (as http transformation)

## System setup
```json
{
  "_id": "email-validation-service",
  "type": "system:microservice",
  "docker": {
    "image": "ohuenno/sesam-email-validator",
    "memory": 32,
    "port": 8080
  },
  "verify_ssl": true
}

```

## Pipe setup

```json
 {
    "type": "http",
    "system": "email-validation-service",
    "url": "Email_Work" <- name of attribute containing email
  }, {
    "type": "http",
    "system": "email-validation-service",
    "url": "Email_Home" <- name of attribute containing email
  }
  ```
  
  transformation result will include new attribute with name equal to checked attribute concatenated with "_validated" and boolean value  
  true for validated address or false otherwise
