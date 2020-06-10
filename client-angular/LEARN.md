
# Understand the code

1. Environment serverUrl

In `environments -> environment.ts`, we have these lines
```
export const environment = {
    ...
    serverUrl: '/api',
    ...
};
```
the serverUrl setting is applied to the app to rewrite any http request using relative path to something like '/api'.
For example, `http.get('/count')` will be rewrite to `http.get('[host]:[port]/api/count')`. Thus all the out-bound requests will be rewrite and handled automatically.

2. proxy to a backend - [Official Angular doc](https://angular.io/guide/build#proxying-to-a-backend-server)
In `proxy.conf.js`, we set up a proxy to redirect any matching traffic to a server. As the documentation explained, this is useful when you run angular as a single server and have another
server running somewhere else with a host and port. In this app, since we use container to host the server under one host address, so this proxy is not used. But helpful to understand it.

The config file is processed in the `npm start` command. (ng serve --proxy-config proxy.conf.js)
