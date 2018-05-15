from datetime import datetime
from flask.json import JSONEncoder
class MySQLStore(object):
    
    def __init__(self, db):
        self._db = db

    def insert(self, new_song):
        with self._db.cursor() as cursor:
            # insert a new time
            sql = "INSERT INTO `songs` (`title`, `artist`, `created_at`, `listened`) VALUES (%s, %s, %s, %s)"
            cursor.execute(sql, (new_song.title, new_song.artist, new_song.created_at, new_song.listened))
            new_song.id = cursor.lastrowid
        self._db.commit()
        return new_song

    def update(self, song_id, listened=False):

        with self._db.cursor() as cursor:
            sql = "UPDATE `songs` SET `listened`=%s WHERE `id`=%s"
            num_rows = cursor.execute(sql, (listened, song_id))
            
        self._db.commit()
        if num_rows == 0:
            return None

        with self._db.cursor() as cursor:
            sql = "SELECT `id`, `title`, `artist`, `created_at`, `listened` FROM `songs` WHERE `id`=%s"
            cursor.execute(sql, (song_id))
            result = cursor.fetchone()
        if result is not None:
            return Song(result['id'], result['title'], result['artist'], result['created_at'], bool(result['listened']))

class Song(object):
    def __init__(self, title, artist, id=None, created_at=datetime.now(), listened=False):
        self._id = id
        self._title = title
        self._artist = artist
        self._created_at = created_at
        self._listened = listened
    
    def __str__(self):
        return ''
    
    def to_dict(self):
        data = {
            'id': self._id,
            'title': self._title,
            'created_at': self._created_at,
            'listened': self._listened
        }
        return data

    @property
    def id(self):
        return self._id
    @id.setter
    def id(self, value):
        self._id = value

    @property
    def title(self):
        return self._title
    @title.setter
    def title(self, value):
        self._title = value

    @property
    def artist(self):
        return self._artist
    @artist.setter
    def artist(self, value):
        self._artist = value

    @property
    def created_at(self):
        return self._created_at
    @created_at.setter
    def created_at(self, value):
        self._created_at = value

    @property
    def listened(self):
        return self._listened
    @listened.setter
    def listened(self, value):
        self._listened = value