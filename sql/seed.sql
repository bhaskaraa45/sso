-- 1. Insert dummy users
INSERT INTO users (name, email, password_hash, is_admin)
VALUES
    ('John Doe', 'john.doe@example.com', '$2a$10$sullllz5Ir2UgpgHsMsEHu91nXhgkK13O2Fx89GVamnjVbz9wyI9q', TRUE), --- password = password123
    ('Jane Smith', 'jane.smith@example.com', '$2a$10$sullllz5Ir2UgpgHsMsEHu91nXhgkK13O2Fx89GVamnjVbz9wyI9q', FALSE), --- password = password123
    ('Alice Johnson', 'alice.johnson@example.com', '$2a$10$sullllz5Ir2UgpgHsMsEHu91nXhgkK13O2Fx89GVamnjVbz9wyI9q', FALSE); --- password = password123

-- 2. Insert dummy clients
INSERT INTO clients (name, email, password_hash, client_id)
VALUES
    ('AwesomeApp', 'contact@awesomeapp.com', '$2a$10$sullllz5Ir2UgpgHsMsEHu91nXhgkK13O2Fx89GVamnjVbz9wyI9q', 'client_abcdef123456'), --- password = password123
    ('SuperService', 'support@superservice.com', '$2a$10$sullllz5Ir2UgpgHsMsEHu91nXhgkK13O2Fx89GVamnjVbz9wyI9q', 'client_ghijkl789012'), --- password = password123
    ('CoolTool', 'admin@cooltool.com', '$2a$10$sullllz5Ir2UgpgHsMsEHu91nXhgkK13O2Fx89GVamnjVbz9wyI9q', 'client_mnopqr345678'); --- password = password123

-- 3. Insert dummy user-client relationships
INSERT INTO user_clients (user_id, client_id)
VALUES
    (1, 1), -- John Doe accessed AwesomeApp
    (1, 2), -- John Doe accessed SuperService
    (2, 1), -- Jane Smith accessed AwesomeApp
    (3, 2), -- Alice Johnson accessed SuperService
    (3, 3); -- Alice Johnson accessed CoolTool