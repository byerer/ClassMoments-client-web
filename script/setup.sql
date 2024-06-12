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

DELIMITER $$

CREATE TRIGGER initCreateAt
    BEFORE INSERT ON users
    FOR EACH ROW
BEGIN
    SET NEW.role := 8889;
END$$

DELIMITER ;
