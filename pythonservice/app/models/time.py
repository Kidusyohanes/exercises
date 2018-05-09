from datetime import datetime


class MySQLStore(object):
    
    def __init__(self, db):
        self._db = db

    def insert(self, new_time):
        pass
    def update(self, time_id, name=None):
        pass

class Time(object):
    def __init__(self, id=None, name=None, start=datetime.now(), end=None):
        pass

    @property
    def id(self):
        return self._id
    @id.setter
    def id(self, value):
        self._id = value

    @property
    def name(self):
        return self._name
    @name.setter
    def name(self, value):
        self._name = value

    @property
    def start(self):
        return self._start
    @start.setter
    def start(self, value):
        self._start = value

    @property
    def end(self):
        return self._end
    @end.setter
    def end(self, value):
        self._end = value