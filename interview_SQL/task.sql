--Сначала создадим структуру бд для библиотеки 
CREATE TABLE books (
  ID INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  genre VARCHAR(255),
  author VARCHAR(255),
  author_birthdate DATE,
  rating DECIMAL (3, 2)
);
CREATE TABLE shops (
  ID INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255)
);
CREATE TABLE book_shop (
  book_id INT,
  shop_id INT,
  PRIMARY KEY (book_id, shop_id),
  FOREIGN KEY (book_id) REFERENCES books(id),
  FOREIGN KEY (shop_id) REFERENCES shops(id)
);
--Заполнение БД для тестов задачи
INSERT INTO books (id, title, genre, author, rating)
VALUES (
    1,
    'Проект «Два тела»',
    'Научпоп',
    'Аса Лунд',
    8.80
  ),
  (
    2,
    'Нормальные люди',
    'Роман',
    'Салли Руни',
    8.60
  ),
  (
    3,
    'Тревожные люди',
    'Роман',
    'Фредрик Бакман',
    8.90
  ),
  (
    4,
    'Маленькие огни повсюду',
    'Роман',
    'Селесте Инг',
    9.00
  ),
  (
    5,
    'Клара и Солнце',
    'Фантастика',
    'Кадзуо Исигуро',
    8.70
  ),
  (
    6,
    'Книжный магазин на берегу',
    'Современная проза',
    'Дженни Колган',
    8.30
  ),
  (
    7,
    'Семья Рой',
    'Нон-фикшн',
    'Джошуа Феррис',
    8.20
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
  (7, 4);
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
SELECT book_shop.shop_id,
  book_shop.book_id,
  shops.id,
  books.title
FROM book_shop
  JOIN shops ON book_shop.shop_id = shops.id
  JOIN books ON book_shop.book_id = books.id
WHERE shops.title != 'Фаланстер';
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
