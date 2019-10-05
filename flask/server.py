import os
from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/', methods=['GET'])
def getAll():
    try:
        with open('../citacoes.json', 'r') as file:
            citations = file.read()

        return jsonify(citations)

    except Exception as err:
        print(err)
        return jsonify({"status": "ERROR"})

if __name__ == "__main__":
    port = int(os.environ.get('PORT', 5000))
    app.run(host='0.0.0.0', port=port, debug=False)
