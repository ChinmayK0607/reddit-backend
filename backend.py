import json

from flask import Flask
from flask import jsonify
from flask import request

app = Flask(__name__)
posts = {
    0:{
      "id": 0,
      "upvotes": 1,
      "title": "My cat is the cutest!",
      "link": "https://i.imgur.com/jseZqNK.jpg",
      "username": "alicia98",
      "comments":[{
        "id": 0,
      "upvotes": 8,
      "text": "Wow, my first Reddit gold!",
      "username": "alicia98",
      }]
    }
}
task_id_ct = 1
comment_id = 1

@app.route("/")
def hello_world():
    return "Welcome to chinmaydit"


# your routes here
@app.route("/posts/")
def get_posts():
    results = {
        "success":True,
        "posts":list(posts.values())
    }
    return json.dumps(results),200

@app.route("/posts/",methods = ["POST"])
def create_post():
    global task_id_ct
    body = json.loads(request.data)
    title = body.get("title","no-title")
    url = body.get("url"," ")
    username = body.get("username","error")
    post1 = {
        "id":task_id_ct,
        "upvotes":0,
        "title": title,
        "link":url,
        "username":username,
        "comments":[]
    }
    posts[task_id_ct]=post1
    task_id_ct +=1
    return json.dumps({"success":True,"posts":post1}),201

@app.route("/posts/<int:task_id>/")
def get_post(task_id):
    post1  = posts.get(task_id)
    if not post1:
        return json.dumps({"success":False,"error":"Task not found"}),404
    else:
        return json.dumps({"success":True,"data":post1}),200

@app.route("/posts/<int:task_id>/",methods = ["DELETE"])
def delete_task(task_id):
   # global task_id_ct
    task1  = posts.get(task_id)
    if not task1:
        return json.dumps({"success":False,"error":"Task not found"}),404
    del posts[task_id]
    #task_id_ct-=1
    return json.dumps({"success":True,"data":task1}),200

@app.route("/posts/<int:task_id>/comments/")
def get_comments(task_id):
    post = posts.get(task_id)
    if not post:
        return json.dumps({"success":False,"error":"Task not found"}),404
    else:
        return json.dumps({"success":True,"comments":list(post["comments"])}),200
@app.route("/posts/<int:task_id>/comments/",methods=["POST"])

def add_comment(task_id):
    global comment_id
    post = posts.get(task_id)
    body = json.loads(request.data)
    title = body.get("title","no-title")
    username = body.get("username","error:please provide a valid username")
    comment = {
        "id":comment_id,
        "upvotes":0,
        "title":title,
        "username":username
    }
    post["comments"].append(comment)
    comment_id+=1
    return json.dumps({"success":True,"posts":post}),201

@app.route("/posts/<int:task_id>/comments/<int:comment_id>/",methods=["POST"])

def update_comment(task_id,comment_id):
    post = posts.get(task_id)
    comment = post["comments"][comment_id]
    body = json.loads(request.data)
    edited_text = body.get("text","error")
    comment1 = {
        "id":comment_id,
        "upvotes":comment["upvotes"],
        "title":edited_text,
        "username":comment["username"]
    }
    post["comments"][comment_id]  = comment1
    return json.dumps({"success":True,"data":post}),200


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000, debug=True)
