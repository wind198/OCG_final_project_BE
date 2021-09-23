SELECT count(id) as total_orders, sum(case when orders.fulfilled_at is not null then orders.total_price else null end) as total_sales, count(case when orders.fulfilled_at is not null then orders.id else null end) as paid_orders, count(case when orders.fulfilled_at is null then orders.id else null end) as unpaid_orders FROM `orders` WHERE (created_at between '2021-09-20 12:16:17.213' and '2021-09-20 12:16:17.213') AND `orders`.`deleted_at` IS NULL