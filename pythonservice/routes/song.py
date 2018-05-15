from flask import request, jsonify
from app import app, song_store
from app.models.song import Song, MySQLStore
from app.routes.errors import bad_request

@app.route('/song', methods=['POST'])
def time_handler():
    data = request.get_json() or {}
    if 'title' not in data or 'artist' not in data:
        return bad_request('must provide valid song to insert')
    new_song = Song(artist=data['artist'], title=data['title'])
    s = song_store.insert(new_song)
    return jsonify(s.to_dict())


@app.route('/song/<id>', methods=['PATCH'])
def specific_time_handler(id):
    data = request.get_json() or {}
    if 'listened' not in data:
        return bad_request('must provide completed to update')     
    else:
        t = song_store.update(id, data['listened'])
    if t is None:
        return bad_request('must provide valid id to update')
    return jsonify(t.to_dict())
