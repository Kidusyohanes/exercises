import os
from dotenv import load_dotenv

# load any environment files
basedir = os.path.abspath(os.path.dirname(__file__))
load_dotenv(os.path.join(basedir, '.env'))

class Config(object):
    MYSQL_HOST = os.environ.get("MYSQL_HOST") or 'localhost'
    MYSQL_PORT = os.environ.get("MYSQL_PORT") or 3306
    MYSQL_DATABASE = os.environ.get("MYSQL_DATABASE") or 'tracking'
    MYSQL_PASSWORD = os.environ.get("MYSQL_PASSWORD") or 'supersecret'
    MYSQL_USER = os.environ.get("MYSQL_USER") or 'root'
