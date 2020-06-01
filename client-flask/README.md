# taylorzh-Flask

Once look into the html side, I realized that flask of ideal for building the website,
but it has the difficulty to do the web app, since the key part - ajax - is in javascript and
python has nothing to do with it - so why not use js based framework to achieve this?


#Gotcha

1. Externally Visible Server
When I first set up the Dockerfile, I use `CMD ["flask", "run"]` to run the flask. However, even if I set the docker `-p` to the right port, my brower still can't reach ths flask app. After 2 hours of research, I figured it out!
We should use `CMD ["flask", "run", "--host=0.0.0.0"]` instead if debug is true.


https://stackoverflow.com/questions/30323224/deploying-a-minimal-flask-app-in-docker-server-connection-issues

```
If you run the server you will notice that the server is only accessible from your own computer, not from any other in the network. This is the default because in debugging mode a user of the application can execute arbitrary Python code on your computer.

If you have the debugger disabled or trust the users on your network, you can make the server publicly available simply by adding --host=0.0.0.0 to the command line:

$ flask run --host=0.0.0.0
This tells your operating system to listen on all public IPs.
```
