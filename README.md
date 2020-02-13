# go-mux-vuejs-buefy-typescript-cloud-template

## Local Project setup
```
npm install
```

### Runs backend and frontend
```
npm start
```

### Compiles and minifies frontend for production
```
npm run build:all
```

### Compiles backend for production
```
npm run build:api
```

### Run your unit tests
```
npm run test:unit
```

### Run your end-to-end tests for frontend
```
npm run test:e2e
```

### Run your end-to-end tests for frontend
```
npm run test:e2e-api
```

### Lints and fixes files
```
npm run lint
```

### Build the container 

```
podman build -t registry.gitlab.com/bobymcbobs/go-mux-vuejs-buefy-typescript-cloud-template .
```

### Deployment

```
kubectl apply -f k8s-manifests
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
