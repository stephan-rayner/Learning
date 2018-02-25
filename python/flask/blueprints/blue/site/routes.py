from flask import Blueprint
mod = Blueprint('site', __name__)

@mod.route('/')
def homepage():
    return '<h1> Welcome to the home page!!!</h1>'
