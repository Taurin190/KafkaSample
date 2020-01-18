var users = [{
    user: "kafka",
    pwd: "kafka",
    roles: [{
        role: "readWrite",
        db: "kafka"
    }]
}];

for (var i = 0, length = users.length; i < length; ++i) {
    db.createUser(users[i]);
}