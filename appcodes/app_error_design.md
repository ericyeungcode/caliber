# API error

- When API service see AppError, it will output error code and message (with ctxvars)
- When API service see builtin error, it will hide the actual error and output "internal error" msg
- API developer should use AppError in function return value
- Low-level(engine, database connection) dev can choose to use AppError

# API output

```json
{"code":1002,"message":"Invalid arguments","vars":{"qty":"-2"}}

{"code":1001,"message":"internal error"}
```


Generally, App/Web/UI should show message in the popup.
While Api user can see the reason details


# Use Pair[*Payload, error]

to facilitate function return handling
