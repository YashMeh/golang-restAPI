create table books
(
    id serial,
    isbn varchar(8),
    title text not null,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT books_pkey PRIMARY KEY (id)
);

insert into books
VALUES(
        "100",
        "book1",
        25.00 
);

insert into books
VALUES(
        "101",
        "book2",
        30.00  
);