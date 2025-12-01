--Сначала создадим структуру бд для библиотеки 
CREATE TABLE books (
  ID INT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  genre VARCHAR(255),
  author VARCHAR(255),
  author_birthdate DATE,
  rating DECIMAL (3, 2)
);
CREATE TABLE shops (ID INT PRIMARY KEY, title VARCHAR(255));
CREATE TABLE book_shop (
  book_id INT,
  shop_id INT,
  PRIMARY KEY (book_id, shop_id),
  FOREIGN KEY (book_id) REFERENCES books(id),
  FOREIGN KEY (shop_id) REFERENCES shops(id)
);
/*
 ---------------------------------------
 
 */
--Написать следующие запросы: 
--Вывести количество книг по жанрам
SELECT genre,
  COUNT(id) AS book_count
FROM books
GROUP BY genre;
--Вывести книги определённого жанра из определённого магазина
SELECT books.title,
  books.genre
FROM books
  JOIN book_shop ON books.id = book_shop.book_id
  JOIN shops ON book_shop.shop_id = shops.id
WHERE books.genre = 'Фантастика'
  AND shops.title = 'Лабиринт';
--Вывести вторую книгу по популярности (по полю rating)
SELECT books.title,
  books.rating
FROM books
ORDER BY rating DESC
LIMIT 1 OFFSET 1;
--Книги, отсутствующие в определённом магазине
SELECT b.id,
  b.title
FROM books b
WHERE b.id NOT IN (
    SELECT book_id
    FROM book_shop bs
      JOIN shops s ON s.id = bs.shop_id
    WHERE s.title = 'Фаланстер'
  );
/*
 ---------------------------------
 
 Вывести книги, авторы которых родились до 1980 года, и количество книг по жанрам
 (Это условие можно было бы проверить, если в таблице книг был бы год рождения автора, 
 но так как его нет в структуре, запрос не применим без изменений.)*/
SELECT b.title,
  b.author,
  b.author_birthdate,
  b.genre,
  g.count_by_genre
FROM (
    SELECT id,
      title,
      author,
      author_birthdate,
      genre
    FROM books
    WHERE author_birthdate < '1980-01-01'
  ) b
  FULL JOIN (
    SELECT genre,
      COUNT(id) AS count_by_genre
    FROM books
    GROUP BY genre
  ) g ON b.genre = g.genre;
--Вывести авторов с количеством их книг, где количество книг больше 3
SELECT b.author,
  COUNT(bs.book_id) AS book_count
FROM books b
  JOIN book_shop bs ON b.id = bs.book_id
GROUP BY b.author
HAVING book_count > 3;
--Вывести список магазинов и средний рейтинг книг в каждом магазине
SELECT s.title,
  AVG(b.rating) AS avg_rating
FROM shops s
  JOIN book_shop bs ON s.id = bs.shop_id
  JOIN books b ON b.id = bs.book_id
GROUP BY s.title;
--Вывести книги, у которых рейтинг ниже 3 и жанр "Драма"
SELECT b.title,
  b.genre,
  b.rating
FROM books b
WHERE genre = 'Драма'
  AND rating < 3;
--Список авторов, у которых есть книги в нескольких жанрах
SELECT b.author,
  COUNT(DISTINCT b.genre) AS genre_count
FROM books b
GROUP BY b.author
HAVING COUNT(DISTINCT b.genre) > 1;
/*
 
 
 --Заполнение БД для тестов задачи
 
 
 */
INSERT INTO books (
    id,
    title,
    genre,
    author,
    author_birthdate,
    rating
  )
VALUES (
    1,
    'Проект «Два тела»',
    'Научпоп',
    'Аса Лунд',
    NULL,
    8.80
  ),
  (
    2,
    'Нормальные люди',
    'Роман',
    'Салли Руни',
    '1991-02-20',
    8.60
  ),
  (
    3,
    'Тревожные люди',
    'Роман',
    'Фредрик Бакман',
    '1981-06-02',
    8.90
  ),
  (
    4,
    'Маленькие огни повсюду',
    'Роман',
    'Селесте Инг',
    '1980-07-30',
    9.00
  ),
  (
    5,
    'Клара и Солнце',
    'Фантастика',
    'Кадзуо Исигуро',
    '1954-11-08',
    8.70
  ),
  (
    6,
    'Книжный магазин на берегу',
    'Современная проза',
    'Дженни Колган',
    '1972-09-14',
    8.30
  ),
  (
    7,
    'Семья Рой',
    'Нон-фикшн',
    'Джошуа Феррис',
    '1974-11-08',
    8.20
  ),
  (
    8,
    'Грустная история',
    'Драма',
    'Автор Драмы',
    '1970-01-01',
    2.50
  ),
  (
    9,
    'Жанровый эксперимент 1',
    'Фантастика',
    'Мульти Автор',
    '1960-05-05',
    7.20
  ),
  (
    10,
    'Жанровый эксперимент 2',
    'Роман',
    'Мульти Автор',
    '1960-05-05',
    8.10
  ),
  (
    11,
    'Исследования хаоса',
    'Научпоп',
    'Академик',
    '1950-02-02',
    7.90
  ),
  (
    12,
    'Тени улиц',
    'Детектив',
    'Академик',
    '1950-02-02',
    8.40
  ),
  (
    13,
    'Серия А',
    'Приключения',
    'Серийный Автор',
    '1975-03-03',
    7.10
  ),
  (
    14,
    'Серия B',
    'Фэнтези',
    'Серийный Автор',
    '1975-03-03',
    7.50
  ),
  (
    15,
    'Серия C',
    'Драма',
    'Серийный Автор',
    '1975-03-03',
    6.80
  ),
  (
    16,
    'Серия D',
    'Хоррор',
    'Серийный Автор',
    '1975-03-03',
    6.40
  );
INSERT INTO shops (id, title)
VALUES (1, 'Читай-город'),
  (2, 'Буквоед'),
  (3, 'Лабиринт'),
  (4, 'Фаланстер');
INSERT INTO book_shop (book_id, shop_id)
VALUES (1, 1),
(1, 3),
  (2, 1),
(2, 2),
  (3, 2),
(3, 3),
(3, 4),
  (4, 1),
(4, 4),
  (5, 2),
(5, 3),
  (6, 1),
(6, 2),
  (7, 4),
  (8, 3),
  (9, 1),
(9, 2),
  (10, 2),
(10, 3),
  (11, 4),
  (12, 1),
  (13, 1),
(13, 2),
(13, 3),
  (14, 2),
(14, 3),
  (15, 3),
(15, 4),
  (16, 4),
(16, 1);
