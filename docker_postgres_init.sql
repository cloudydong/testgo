CREATE TABLE corp(
    id bigint NOT NULL,
    name text COLLATE pg_catalog."default",
    CONSTRAINT corp_pkey PRIMARY KEY (id)
);

INSERT INTO corp(id, name) VALUES
(005930, '삼성전자'),
(139480, '이마트'),
(004170, '신세계');
