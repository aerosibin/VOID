-- Add unique constraint to prevent duplicate saved items
ALTER TABLE public.saved_items 
ADD CONSTRAINT unique_user_product UNIQUE (user_id, product_name);