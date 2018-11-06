from flask import Flask, request

from pymongo import MongoClient
DB_URL = 'mongodb://127.0.0.1:27017/'
client = MongoClient(DB_URL)

app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello World!"


todos_db = client["todos"]
todos_collection = todos_db.todos

todos_collection.insert_one({"title": "My Title", "description": "This is a description", "done": False})


print( [ x for x in todos_collection.find({}) ])

