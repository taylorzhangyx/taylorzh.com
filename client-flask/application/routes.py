from application import app
from flask import (
    render_template,
    request,
    Response,
    json,
    jsonify,
    redirect,
    flash,
    url_for,
)
import requests


@app.route("/flask/", methods=["GET"])
@app.route("/flask/index", methods=["GET"])
@app.route("/flask/home", methods=["GET"])
def index():
    print("wow!")
    return render_template("index.html")


@app.route("/flask/count")
def count():
    res = requests.get("http://server-gin:8080")
    obj = json.loads(res.text)
    return render_template("count.html", count=obj["count"])
