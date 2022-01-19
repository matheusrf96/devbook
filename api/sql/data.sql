INSERT INTO users (name, username, email, password)
VALUES
('Matheus', 'matheus', 'matheus@devbook.com', '$2a$04$2DMts1cJsSmnrbMcV9m.yOHeSXy890WAQanBXfHddG6tKrOlISzVK')
, ('Teste', 'teste', 'teste@devbook.com', '$2a$04$2DMts1cJsSmnrbMcV9m.yOHeSXy890WAQanBXfHddG6tKrOlISzVK')
, ('Corinthians', 'corinthians', 'corinthians@corinthians.com.br', '$2a$04$2DMts1cJsSmnrbMcV9m.yOHeSXy890WAQanBXfHddG6tKrOlISzVK');

INSERT INTO user_followers (user_id, follower_id)
VALUES
(1, 3)
, (1, 2)
, (3, 1)
, (3, 2);