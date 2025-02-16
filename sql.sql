SELECT
    p.id,
    p.name,
    p.price,
    p.category_id,
    p.created_at,
    p.updated_at,

    JSON_BUILD_OBJECT(
    'id', c.id,
    'title', c.title,
    'parent_id', c.parent_id,
    'updated_at', c.updated_at,
    'created_at', c.created_at
    ) AS category
FROM product as p
LEFT JOIN category AS c ON c.id = p.category_id
WHERE p.id = 'c60cabb8-a7e9-48a7-b8c2-6ae7569c03fb';


WITH product_branches AS (
    SELECT
        bp.product_id,
        JSON_AGG(
            JSON_BUILD_OBJECT(
                'id', b.id,
                'name', b.name,
                'address', b.address,
                'phone_number', b.phone_number,
                'created_at', b.created_at,
                'updated_at', b.updated_at
            )
        ) AS branches
    FROM branch AS b
    JOIN branch_product_relation AS bp ON bp.branch_id = b.id
    WHERE bp.product_id = 'c60cabb8-a7e9-48a7-b8c2-6ae7569c03fb'
    GROUP BY bp.product_id
)
SELECT
    p.id,
    p.name,
    p.price,
    p.category_id,
    p.created_at,
    p.updated_at,
    COALESCE(pb.branches, '[]'::json) as branches,

    JSON_BUILD_OBJECT(
    'id', c.id,
    'title', c.title,
    'parent_id', c.parent_id,
    'updated_at', c.updated_at,
    'created_at', c.created_at
    ) AS category
FROM product AS p 
JOIN product_branches AS pb ON pb.product_id = p.id
LEFT JOIN category AS c ON c.id = p.category_id
WHERE p.id = 'c60cabb8-a7e9-48a7-b8c2-6ae7569c03fb';