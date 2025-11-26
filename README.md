# my fullstack app

This is fulltsack project related Auth using golang as backend and nuxt as frontend.
Backend using golang go-chi for http handler and frontend using latest nuxt version 4 fot better LTS.
Using openapi.yaml to define api specification on backend and then generate api client on frontend.
Using module auth based on api specification. Later we can create v2 inside those module if needed.
For persistent data using sqlite for lightweight use and data already seeded during run.
writing backend unit test on repository, usecase, and service also using mockgen.
writing frontend unit test on page level using @nuxt/test-utils.

list of tools version of my machine:

```bash
go version go1.24.0 darwin/arm64
node v24.11.1
```

please use gvm and nvm to install the version

install all related requirements:

```bash
gvm install go1.24
gvm use go1.24
```

incase need to reload env after install new Go:

```bash
source ~/.zshrc
```

based on your terminal like .bashrc

```bash
nvm install v24.11.1
nvm use v24.11.1
npm install -g pnpm
```

Run backend server on local:

```bash
cd backend
cp env.sample .env
make tool-openapi
make openapi-gen
make dep
make gen-secret
make run
```

Run backend server on production build:

```bash
make build
./bin/mygolangapp
```

Run frontend on local:

```bash
cp env.sample .env
pnpm install
pnpm openapigen
pnpm dev
```

Run frontend on production build:

```bash
pnpm build
pnpm preview
```

To checking openapi documentations, you can visit this url after backend running.

```bash
http://localhost:8080/docs/
```

Register by visiting:

```bash
http://localhost:3000/register
```

Login by visiting:

```bash
http://localhost:3000/login
```

sample credentials:

```bash
{
    "email": "test@test.com",
    "password": "password"
}
```

Visit dashboard:

```bash
http://localhost:3000/dashboard
```


evidences [video](https://drive.google.com/file/d/1wlLAGCXSUQO50eZW8FeTXmEOh4ec5zbU/view?usp=sharing)
see backend [README.md](backend/README.md)
see frontend [README.md](frontend/README.md)
