TRUNCATE TABLE micro.public.user;
INSERT INTO micro.public.user (id, username, nick_name, email, password, authority)
VALUES (1, 'admin', 'admin', 'admin@gmail.com', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'SYS_ADMIN');