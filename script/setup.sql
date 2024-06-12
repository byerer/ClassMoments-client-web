create table classes
(
    class_id          bigint unsigned auto_increment
        primary key,
    class_name        longtext null,
    school_name       longtext null,
    class_description longtext null
);

create table comments
(
    comment_id bigint unsigned auto_increment
        primary key,
    user_id    bigint unsigned null,
    moment_id  bigint unsigned null,
    content    longtext        null,
    created_at datetime(3)     null
);

create table likes
(
    like_id    bigint unsigned auto_increment
        primary key,
    liker_id   bigint unsigned null,
    likee_id   bigint unsigned null,
    moment_id  bigint unsigned null,
    created_at datetime(3)     null
);

create table media
(
    media_id   bigint unsigned auto_increment
        primary key,
    moment_id  bigint unsigned null,
    media_type longtext        null,
    url        longtext        null,
    created_at longtext        null
);

create table moments
(
    moment_id  bigint unsigned auto_increment
        primary key,
    user_id    bigint unsigned null,
    class_id   bigint unsigned null,
    role       bigint unsigned null,
    title      longtext        null,
    content    longtext        null,
    image      longtext        null,
    created_at datetime(3)     null
);

create table notifications
(
    notification_id bigint unsigned auto_increment
        primary key,
    user_id         bigint unsigned null,
    moment_id       bigint unsigned null,
    is_read         bigint          null,
    type            longtext        null,
    create_at       datetime(3)     null
);

create table parent_rel_children
(
    parent_id bigint unsigned null,
    child_id  bigint unsigned null
);

create table schools
(
    school_id   bigint unsigned auto_increment
        primary key,
    school_name longtext null
);

create table teachers
(
    teacher_id     bigint unsigned null,
    subject_taught longtext        null
);

create table users
(
    user_id     bigint unsigned auto_increment
        primary key,
    role        bigint unsigned null,
    sex         bigint unsigned null,
    class_id    bigint unsigned null,
    school_name longtext        null,
    name        longtext        null,
    username    varchar(191)    null,
    password    longtext        null,
    created_at  datetime(3)     null,
    updated_at  datetime(3)     null,
    constraint uni_users_username
        unique (username)
);

-- 创建视图

-- 显示每个用户的动态、评论和点赞数量
CREATE VIEW view_user_activity AS
SELECT
    u.user_id, u.username,
    COUNT(DISTINCT m.moment_id) AS moment_count,
    COUNT(DISTINCT c.comment_id) AS comment_count,
    COUNT(DISTINCT l.like_id) AS like_count
FROM
    users u
        LEFT JOIN moments m ON u.user_id = m.user_id
        LEFT JOIN comments c ON u.user_id = c.user_id
        LEFT JOIN likes l ON u.user_id = l.liker_id
GROUP BY
    u.user_id, u.username;

-- 显示每个动态的评论和点赞数量
CREATE VIEW view_moment_engagement AS
SELECT
    m.moment_id, m.title, m.content, m.created_at,
    COUNT(DISTINCT c.comment_id) AS comment_count,
    COUNT(DISTINCT l.like_id) AS like_count
FROM
    moments m
        LEFT JOIN comments c ON m.moment_id = c.moment_id
        LEFT JOIN likes l ON m.moment_id = l.moment_id
GROUP BY
    m.moment_id, m.title, m.content, m.created_at;

-- 显示每个用户所属的班级及学校详细信息
CREATE VIEW view_user_class_details AS
SELECT
    u.user_id, u.username, u.name,
    c.class_id, c.class_name, c.school_name AS class_school_name,
    s.school_id, s.school_name AS school_name
FROM
    users u
        LEFT JOIN classes c ON u.class_id = c.class_id
        LEFT JOIN schools s ON u.school_name = s.school_name;

-- 显示每个动态的详细信息，包括发布者、媒体和评论
CREATE VIEW view_moment_full_details AS
SELECT
    m.moment_id, m.title, m.content, m.image, m.created_at,
    u.user_id, u.username,
    med.media_id, med.media_type, med.url,
    c.comment_id, c.content AS comment_content, c.created_at AS comment_created_at, cu.username AS comment_user
FROM
    moments m
        LEFT JOIN users u ON m.user_id = u.user_id
        LEFT JOIN media med ON m.moment_id = med.moment_id
        LEFT JOIN comments c ON m.moment_id = c.moment_id
        LEFT JOIN users cu ON c.user_id = cu.user_id;

-- 显示每个用户的未读通知数量
CREATE VIEW view_notifications_overview AS
SELECT
    u.user_id, u.username,
    COUNT(n.notification_id) AS unread_notifications
FROM
    users u
        LEFT JOIN notifications n ON u.user_id = n.user_id AND n.is_read = 0
GROUP BY
    u.user_id, u.username;

-- 显示每个用户的最近活动，包括动态、评论和点赞
CREATE VIEW view_user_recent_activity AS
SELECT
    u.user_id, u.username,
    m.moment_id, m.title AS moment_title, m.created_at AS moment_created_at,
    c.comment_id, c.content AS comment_content, c.created_at AS comment_created_at,
    l.like_id, l.created_at AS like_created_at
FROM
    users u
        LEFT JOIN moments m ON u.user_id = m.user_id
        LEFT JOIN comments c ON u.user_id = c.user_id
        LEFT JOIN likes l ON u.user_id = l.liker_id
ORDER BY
    COALESCE(m.created_at, c.created_at, l.created_at) DESC;

-- 显示每个班级的动态及其发布者信息
CREATE VIEW view_class_moments AS
SELECT
    c.class_id, c.class_name, c.school_name,
    m.moment_id, m.title, m.content, m.created_at,
    u.user_id, u.username
FROM
    classes c
        LEFT JOIN moments m ON c.class_id = m.class_id
        LEFT JOIN users u ON m.user_id = u.user_id;

-- 显示每条评论的详细信息，包括评论者和被评论的动态信息
CREATE VIEW view_comment_details AS
SELECT
    c.comment_id, c.content, c.created_at,
    u.user_id, u.username,
    m.moment_id, m.title AS moment_title, m.content AS moment_content
FROM
    comments c
        LEFT JOIN users u ON c.user_id = u.user_id
        LEFT JOIN moments m ON c.moment_id = m.moment_id;

-- 显示每个用户的家长和孩子信息
CREATE VIEW view_user_relationships AS
SELECT
    p.user_id AS parent_id, p.username AS parent_username,
    c.user_id AS child_id, c.username AS child_username
FROM
    parent_rel_children prc
        LEFT JOIN users p ON prc.parent_id = p.user_id
        LEFT JOIN users c ON prc.child_id = c.user_id;

-- 显示每个学校的班级数量和学生数量
CREATE VIEW view_school_overview AS
SELECT
    s.school_id, s.school_name,
    COUNT(DISTINCT c.class_id) AS class_count,
    COUNT(DISTINCT u.user_id) AS student_count
FROM
    schools s
        LEFT JOIN classes c ON s.school_name = c.school_name
        LEFT JOIN users u ON c.class_id = u.class_id
GROUP BY
    s.school_id, s.school_name;


DELIMITER $$

CREATE TRIGGER initCreateAt
    BEFORE INSERT ON users
    FOR EACH ROW
BEGIN
    SET NEW.role := 8889;
END$$

CREATE TRIGGER before_insert_moment_check_user_existence
    BEFORE INSERT ON moments
    FOR EACH ROW
BEGIN
    DECLARE user_exists INT;
    SELECT COUNT(*) INTO user_exists FROM users WHERE user_id = NEW.user_id;
    IF user_exists = 0 THEN
    SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'User does not exist';
END IF;
END$$



DELIMITER ;
