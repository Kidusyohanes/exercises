from flask import Flask
# from flask.json import JSONEncoder
import pymysql.cursors
from config import Config
from app.models.song import MySQLStore

app = Flask(__name__)
app.config.from_object(Config)

db = pymysql.connect(host=app.config['MYSQL_HOST'],
                     port=int(app.config['MYSQL_PORT']),
                     user=app.config['MYSQL_USER'],
                     password=app.config['MYSQL_PASSWORD'],
                     db=app.config['MYSQL_DATABASE'],
                     charset='utf8',
                     cursorclass=pymysql.cursors.DictCursor)

# create a new time store for use in the routeimp
song_store = MySQLStore(db)

from app.routes import rdfa