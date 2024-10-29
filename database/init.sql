-- init.sql

-- Criação da tabela orders com campos id, item, quantity, e price
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    item VARCHAR(100) NOT NULL,
    quantity INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

-- Exemplo de dados iniciais (opcional)
INSERT INTO orders (item, quantity, price) VALUES ('Laptop', 2, 1500.00);
INSERT INTO orders (item, quantity, price) VALUES ('Mouse', 10, 20.50);
INSERT INTO orders (item, quantity, price) VALUES ('Keyboard', 5, 45.00);
