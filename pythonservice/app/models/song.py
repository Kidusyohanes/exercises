from datetime import datetime


class MySQLStore(object):
    
    def __init__(self, db):
        self._db = db

    def insert(self, new_time):
        pass
    def update(self, time_id, name=None):
        pass

class Song(object):
    def __init__(self, title, artist, id=None, created_at=datetime.now(), listened=False):
        pass
        
    def to_dict(self):
       pass

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