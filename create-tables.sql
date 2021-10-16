DROP TABLE IF EXISTS song;
CREATE TABLE song (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO song
    (title, artist)
VALUES
    ('Blow With the Fires', 'Between August and December'),
    ('Within', 'Daft Punk'),
    ('Paralyzed', 'Sueco'),
    ('Bol', 'Buerak');