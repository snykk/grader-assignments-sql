-- TODO: answer here
SELECT o.id AS order_id, u.fullname, u.email, o.product_name, o.unit_price, o.quantity, o.order_date
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE u.status = 'active' AND o.unit_price * o.quantity > 500000 OR o.quantity > 20
