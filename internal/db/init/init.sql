CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL,
    start_date VARCHAR(255) NOT NULL,
    end_date VARCHAR(255)
);

INSERT INTO subscriptions (id, service_name, price, user_id, start_date, end_date) VALUES
    ('5d0f12ea-e78d-46f7-bcd5-df6f9ad1c501', 'Yandex Plus', 400, '60601fee-2bf1-4721-ae6f-7636e79a0cba', '2025-07', '2025-09'),
    ('4e99d8f8-6b60-4e68-9c29-1a2f625f3e32', 'Spotify', 300, '60601fee-2bf1-4721-ae6f-7636e79a0cba', '2025-07', NULL),
    ('c65c4be3-4869-4b1c-8780-3b8ad51d20b0', 'Netflix', 500, 'e9c3a059-089d-4c5b-b0cc-17b5a78bb6ef', '2025-07', '2025-08'),
    ('69d51763-479e-4a35-82cf-2337bb4937fe', 'Yandex Plus', 400, '8c431d9e-107f-4a04-b325-d232dd86a9e9', '2025-08', NULL),
    ('0d5b0f55-2f71-433d-a3cd-02099a2db0cf', 'Netflix', 500, '60601fee-2bf1-4721-ae6f-7636e79a0cba', '2025-08', NULL),
    ('1dfe4779-99a9-4645-96d4-032dfb107670', 'Spotify', 300, '8c431d9e-107f-4a04-b325-d232dd86a9e9', '2025-08', '2025-11'),
    ('1a4b213f-f5ad-402b-8a18-dc650ffeb9a7', 'Spotify', 300, '60601fee-2bf1-4721-ae6f-7636e79a0cba', '2025-09', '2025-10');
