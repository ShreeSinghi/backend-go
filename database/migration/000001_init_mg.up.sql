CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  `hash` VARCHAR(255) NOT NULL,
  `admin` BOOLEAN DEFAULT false,
  requested BOOLEAN DEFAULT false
);

CREATE TABLE books (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  quantity INT NOT NULL DEFAULT 1
);


CREATE TABLE cookies (
  id INT AUTO_INCREMENT PRIMARY KEY,
  userId INT,
  sessionId VARCHAR(255),
  FOREIGN KEY (userId) REFERENCES users(id)
);

CREATE TABLE requests (
  id INT AUTO_INCREMENT PRIMARY KEY,
  bookId INT,
  userId INT,
  `state` ENUM('owned', 'outrequested', 'inrequested') DEFAULT 'outrequested',
  FOREIGN KEY (bookId) REFERENCES books(id),
  FOREIGN KEY (userId) REFERENCES users(id)
);

