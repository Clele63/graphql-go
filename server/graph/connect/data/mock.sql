SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE `task_assignees`;
TRUNCATE TABLE `comments`;
TRUNCATE TABLE `tasks`;
TRUNCATE TABLE `columns`;
TRUNCATE TABLE `boards`;
TRUNCATE TABLE `users`;
SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO `users` (id, name, email, password) VALUES
('u1', 'Alice', 'alice@example.com', '{{pwd_password}}'),
('u2', 'Bob', 'bob@example.com', '{{pwd_password}}'),
('u3', 'Carol', 'carol@example.com', '{{pwd_password}}');

INSERT INTO `boards` (id, name) VALUES
('b1', 'Board - Promo');

INSERT INTO `columns` (id, name, `order`, board_id) VALUES
('c1', 'Todo', 1, 'b1'),
('c2', 'Doing', 2, 'b1'),
('c3', 'Done', 3, 'b1');

INSERT INTO `tasks` (id, title, description, column_id) VALUES
('t1', 'Préparer slides', 'Séance 5', 'c1'),
('t2', 'Corriger TP', NULL, 'c1'),
('t3', 'Démo subscription', 'Live coding', 'c2');

INSERT INTO `task_assignees` (task_id, user_id) VALUES
('t1', 'u1'),
('t1', 'u2'),
('t2', 'u2'),
('t3', 'u3');

INSERT INTO `comments` (id, content, author_id, task_id, created_at) VALUES
('cm1', "N'oublie pas les animations!", 'u2', 't1', '2025-11-02T10:30:00'),
('cm2', "Je peux t'aider si besoin", 'u3', 't1', '2025-11-02T14:20:00');