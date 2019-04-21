
from flask import Flask, render_template
app = Flask(__name__)

@app.route('/')
def accueil():
    return render_template('accueil.html')

app = Flask(__name__)

@app.route('/')
def index():
    return "<a href='http://127.0.0.1:5000/user'>clic here</a>"

@app.route('/user/')
def user():
    name = "Connard"
    tel = "01 23 45 67 89"
    return "Name: {} --- Tel: {}".format(name, tel)

@app.route('/discussion/page/<num_page>')
def mon_chat(num_page):
    num_page = int(num_page)
    premier_msg = 1 + 50 * (num_page - 1)
    dernier_msg = premier_msg + 50
    return 'affichage des messages {} à {}'.format(premier_msg, dernier_msg)

@app.route('/template/')
def accueil():
    mots = ["bonjour", "à", "toi,", "visiteur."]
    return render_template('template.html', titre="Bienvenue !", mots=mots)

if __name__ == '__main__':
    app.run(debug=True)
