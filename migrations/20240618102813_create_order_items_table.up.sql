CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT,
    book_id INT,
    quantity INT,
    price DECIMAL(15, 2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(36),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(36),
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    deleted_by VARCHAR(36),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);