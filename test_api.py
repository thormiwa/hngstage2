import requests

BASE_URL = "https://hngstage2-39oa.onrender.com/api"

def create_person(name):
    resp = requests.post(BASE_URL, json={"name": name})
    return resp

def get_person(id):
    resp = requests.get(BASE_URL + f"/{id}")
    return resp

def update_person(id, name):
    resp = requests.put(BASE_URL + f"/{id}", json={"name": name})
    return resp

def delete_person(id):
    resp = requests.delete(BASE_URL + f"/{id}")
    return resp

def test_person():
    name = "Tomiwa"
    new_name = "Jibola"

    # Create a person
    resp = create_person(name)
    print(f"Create person: {resp.json()}")

    # Create person with same name
    resp = create_person(name)
    print(f"Create person: {resp.json()}")

    # Get person
    resp = get_person(1)
    print(f"Get person: {resp.json()}")

    # Update person
    resp = update_person(1, new_name)
    print(f"Update person: {resp.json()}")

    # Delete person
    resp = delete_person(1)
    print(f"Delete person: {resp.json()}")

test_person()