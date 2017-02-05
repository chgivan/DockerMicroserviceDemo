from flask import Flask, jsonify, abort, request, url_for
from redis import Redis

app = Flask(__name__)
redis = Redis(host='redis', port=6379)

IdTemplate = "likes:%s"

@app.route('/likes/<string:likeID>', methods=['GET'])
def get_like(likeID):
	if redis.exists( IdTemplate % likeID):
		count = int(redis.get(IdTemplate % likeID))
		return jsonify( new_like( likeID, count ) ), 200
		
	redis.set(IdTemplate % likeID, 0)
	return jsonify( new_like( likeID, 0 ) ), 200


def new_like(id, count):
	return {
		'id': id,
		'count': count,
	}

@app.route('/likes/<string:likeID>', methods=['POST'])
def sumbit_like(likeID):
	if redis.exists( IdTemplate % likeID):
		count = redis.incr(IdTemplate % likeID)
		return jsonify( new_like( likeID, count ) ), 200 
	return jsonify({'error':"Like ID %s not found" % likeID}), 404


@app.errorhandler(404)
def not_found(error):
	return jsonify({'error':"Not found"}), 404


@app.route('/')
def index():
	return "Welcome"  

if __name__=="__main__":
	app.run(host='0.0.0.0', port=80)

