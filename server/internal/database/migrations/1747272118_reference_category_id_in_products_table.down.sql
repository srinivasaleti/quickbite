DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'products'
          AND column_name = 'category_id'
    ) THEN
        ALTER TABLE products
        DROP COLUMN category_id;
    END IF;
END $$;
