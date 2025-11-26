# mygolangapp

Scaffold demonstrating entity / repository / usecase / service separation.

generate openapi:

```bash
make openapi-gen
```

generate JWT_SECRET:

```bash
make gen-secret
```

Run server:

```bash
make tool-openapi
make openapi-gen
make dep
make gen-secret
make run
```

Format and lint:

```bash
make check
```

Run unit test:

```bash
make test
```

Run unit test with coverage report visualization:

```bash
make coverage-html
```

API:

- POST /v1/auth/register {email,password,confirmPassword}
- POST /v1/auth/login {email,password}
