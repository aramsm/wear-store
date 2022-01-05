# wear-store
A digital store of clothes, build in GO

# Endpoints

| URL | Params | Description |
|-----|--------|-------------|
| http://localhost:3000/api/v1/employees/ | nil | Lists all `employees` from all `stores` |
| http://localhost:3000/api/v1/employees/ | id::string | Shows an `employee` with the `id` sent |
| http://localhost:3000/api/v1/addresses/ | nil | Lists all `addresses` from all `stores` |
| http://localhost:3000/api/v1/addresses/ | street::string | Shows an `address` with the `street` sent |
| http://localhost:3000/api/v1/stores/ | nil | Lists all `stores` |
| http://localhost:3000/api/v1/stores/ | id::string | Shows a `store` with the `id` sent |