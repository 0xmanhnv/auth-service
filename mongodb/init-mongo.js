db = db.getSiblingDB('auth_db'); // Chọn database

db.users.insertMany([
  { username: 'user1', password: 'password1' },
  { username: 'user2', password: 'password2' }
]);
