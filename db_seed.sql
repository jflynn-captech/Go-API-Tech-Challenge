DROP TABLE IF EXISTS person_course;
DROP TABLE IF EXISTS course;
DROP TABLE IF EXISTS person;

-- person
CREATE TABLE person
(
    id         SERIAL PRIMARY KEY,
    guid       VARCHAR(55)                                   NOT NULL,
    first_name TEXT                                          NOT NULL,
    last_name  TEXT                                          NOT NULL,
    email      TEXT                                          NOT NULL,
    type       TEXT CHECK (type IN ('professor', 'student')) NOT NULL,
    age        INTEGER                                       NOT NULL
);

INSERT INTO person (guid, first_name, last_name, email, type, age)
VALUES ('abcd', 'Steve', 'Jobs', 'sjobs@test.com', 'professor', 56),
       ('efgh', 'Jeff', 'Bezos', 'jbezos@test.com', 'professor', 60),
       ('ijkl', 'Larry', 'Page', 'lpage@test.com', 'student', 51),
       ('mnop', 'Bill', 'Gates', 'bgates@test.com', 'student', 67),
       ('qrst', 'Elon', 'Musk',  'emusk@test.com', 'student', 52),
       ('uvwx', 'John', 'Flynn', 'jflynn@test.com', 'student', 52),
       ('yzab', 'John', 'Flynn',  'jflynn2@test.com', 'student', 52);

-- course
CREATE TABLE course
(
    id   SERIAL PRIMARY KEY,
    guid       VARCHAR(55) NOT NULL,
    name TEXT NOT NULL
);

INSERT INTO course (guid, name)
VALUES ('123a', 'Programming'),
       ('456b', 'Databases'),
       ('789c', 'UI Design');

-- person_course
CREATE TABLE person_course
(
    person_id INTEGER NOT NULL,
    course_id INTEGER NOT NULL,
    PRIMARY KEY (person_id, course_id),
    FOREIGN KEY (person_id) REFERENCES person (id),
    FOREIGN KEY (course_id) REFERENCES course (id)
);

INSERT INTO person_course (person_id, course_id)
VALUES (1, 1),
       (1, 2),
       (1, 3),
       (2, 1),
       (2, 2),
       (2, 3),
       (3, 1),
       (3, 2),
       (3, 3),
       (4, 1),
       (4, 2),
       (4, 3),
       (5, 1),
       (5, 2),
       (5, 3);