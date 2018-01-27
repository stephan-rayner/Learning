import shelve

db_dir = "cheese.data"
"""
with shelve.open(db_dir) as db:
    db['item1'] = 1
    db['item2'] = 'default'
    db['item3'] = 2
"""
with shelve.open(db_dir) as db:
    for key in db:
        print(key, db[key])
