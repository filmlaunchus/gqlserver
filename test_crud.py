import requests

url = 'http://localhost:8080/gql'

# Create a user object

create = """mutation { createUser(email: "fakeemail", username: "fakename") { id username email } }"""
resp = requests.post(url, params={'query':create})
data = resp.json()["data"]
id   = str(data["createUser"]["id"])
print("created: ", str(data["createUser"]))

read = 'query { user(id: "'+id+'") { id username email } }'
resp = requests.post(url, params={'query':read})
data = resp.json()["data"]
print("read: ", str(data))

updateWhole = 'mutation { updateUser(id: "'+id+'", email: "fakeemail2", username: "fakename2") { id username email } }'
resp = requests.post(url, params={'query':updateWhole})
data = resp.json()["data"]
print("updateWhole: ", str(data))

read = 'query { user(id: "'+id+'") { id username email } }'
resp = requests.post(url, params={'query':read})
data = resp.json()["data"]
print("read: ", str(data))

updateEmail = 'mutation { updateUser(id: "'+id+'", email: "justemail") { id username email } }'
resp = requests.post(url, params={'query':updateEmail})
data = resp.json()["data"]
print("updateEmail: ", str(data))

read = 'query { user(id: "'+id+'") { id username email } }'
resp = requests.post(url, params={'query':read})
data = resp.json()["data"]
print("read: ", str(data))

delete = 'mutation { deleteUser(id: "'+id+'") { id username email } }'
resp = requests.post(url, params={'query':delete})
data = resp.json()["data"]
print("delete: ", str(data))

read = 'query { user(id: "'+id+'") { id username email } }'
resp = requests.post(url, params={'query':read})
data = resp.json()["data"]
print("read: ", str(data))
