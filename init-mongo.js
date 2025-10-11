db.createUser({
    user: "appuser",
    pwd: "apppassword",
    roles: [
      {
        role: "readWrite",
        db: "appdb"
      }
    ]
  });
  
  db.createCollection("users");
  db.createCollection("products");