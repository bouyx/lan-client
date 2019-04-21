
from flask import Flask
app = Flask(__name__)

@app.route('/home')
def index():
    return "<a href='http://127.0.0.1:5000/test'>clic here</a> "

@app.route('/test')
def test():
    return "Hello connard !"

if __name__ == '__main__':
    app.run(debug=True)
